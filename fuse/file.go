// +build !windows

package fuse

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

var (
	// ErrNotCached is returned in offline mode when we don't have a file
	ErrNotCached = errors.New("content is not cached and need to be downloaded")

	// ErrTooManyWriters is returned when writers counter is about to be overfilled
	ErrTooManyWriters = errors.New("too many writers for the file")
)

// File is a file inside a directory.
type File struct {
	mu   sync.Mutex // used during handle creation
	path string
	m    *Mount
	hd   *Handle
}

// Attr is called to get the stat(2) attributes of a file.
func (fi *File) Attr(ctx context.Context, attr *fuse.Attr) error {
	defer logPanic("file: attr")

	info, err := fi.m.fs.Stat(fi.path)
	if err != nil {
		return err
	}
	debugLog("exec file attr: %v", fi.path)

	var filePerm os.FileMode = 0640
	attr.Mode = filePerm
	if fi.m.options.Offline {
		isCached, err := fi.m.fs.IsCached(fi.path)
		if err != nil || !isCached {
			if err != nil {
				log.Errorf("IsCached failed for %s with error : %v", fi.path, err)
			}
			// Uncached file will be shown as symlink
			// We cannot read them in Offline mode,
			// but we can delete such link and overwrite its content
			attr.Mode = os.ModeSymlink | filePerm
		}
	}
	if fi.hd != nil && fi.hd.writers > 0 {
		attr.Size = uint64(len(fi.hd.data))
	} else {
		attr.Size = info.Size
	}
	attr.Mtime = info.ModTime
	attr.Inode = info.Inode

	// Act like the file is owned by the user of the brig process.
	attr.Uid = uint32(os.Getuid())
	attr.Gid = uint32(os.Getgid())

	// tools like `du` rely on this for size calculation
	// (assuming every fs block takes actual storage, but we only emulate this
	// here for compatibility; see man 2 stat for the why for "512")
	attr.BlockSize = 4096
	attr.Blocks = info.Size / 512
	if info.Size%uint64(512) > 0 {
		attr.Blocks++
	}

	return nil
}

// Open is called to get an opened handle of a file, suitable for reading and writing.
func (fi *File) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	defer logPanic("file: open")
	debugLog("fuse-open: %s", fi.path)

	// Check if the file is actually available locally.
	if fi.m.options.Offline {
		isCached, err := fi.m.fs.IsCached(fi.path)
		if err != nil {
			return nil, errorize("file-is-cached", err)
		}

		if !isCached {
			return nil, errorize("file-not-cached", ErrNotCached)
		}
	}

	fd, err := fi.m.fs.Open(fi.path)
	if err != nil {
		return nil, errorize("file-open", err)
	}

	fi.mu.Lock()
	if fi.hd == nil {
		hd := Handle{fd: fd, m: fi.m, writers: 0, wasModified: false, currentFileReadOffset: -1}
		fi.hd = &hd
	}
	fi.hd.fd = fd
	fi.mu.Unlock()
	if req.Flags.IsReadOnly() {
		// we don't need to track read-only handles
		// and no need to set handle `data`
		return fi.hd, nil
	}

	// for writers we need to copy file data to the handle `data`
	fi.hd.mu.Lock()
	defer fi.hd.mu.Unlock()
	if fi.hd.writers == 0 {
		err = fi.hd.loadData(fi.path)
		if err != nil {
			return nil, errorize("file-open-loadData", err)
		}
	}
	if fi.hd.writers == (^uint(0)) { // checks against writers overflow
		return nil, errorize(fi.path, ErrTooManyWriters)
	}
	fi.hd.writers++
	return fi.hd, nil
}

// Setattr is called once an attribute of a file changes.
// Most importantly, size changes are reported here, e.g. after truncating a
// file, the size change is noticed here before Open() is called.
func (fi *File) Setattr(ctx context.Context, req *fuse.SetattrRequest, resp *fuse.SetattrResponse) error {
	defer logPanic("file: setattr")

	// This is called when any attribute of the file changes,
	// most importantly the file size. For example it is called when truncating
	// the file to zero bytes with a size change of `0`.
	debugLog("exec file setattr")
	switch {
	case req.Valid&fuse.SetattrSize != 0:
		if err := fi.hd.truncate(req.Size); err != nil {
			return errorize("file-setattr-size", err)
		}
	case req.Valid&fuse.SetattrMtime != 0:
		if err := fi.m.fs.Touch(fi.path); err != nil {
			return errorize("file-setattr-mtime", err)
		}
	}

	return nil
}

// Fsync is called when any open buffers need to be written to disk.
// Currently, fsync is completely ignored.
func (fi *File) Fsync(ctx context.Context, req *fuse.FsyncRequest) error {
	defer logPanic("file: fsync")

	debugLog("exec file fsync")
	return nil
}

// Getxattr is called to get a single xattr (extended attribute) of a file.
func (fi *File) Getxattr(ctx context.Context, req *fuse.GetxattrRequest, resp *fuse.GetxattrResponse) error {
	defer logPanic("file: getxattr")

	debugLog("exec file getxattr: %v: %v", fi.path, req.Name)
	xattrs, err := getXattr(fi.m.fs, req.Name, fi.path, req.Size)
	if err != nil {
		return err
	}

	resp.Xattr = xattrs
	return nil
}

// Listxattr is called to list all xattrs of this file.
func (fi *File) Listxattr(ctx context.Context, req *fuse.ListxattrRequest, resp *fuse.ListxattrResponse) error {
	defer logPanic("file: listxattr")

	debugLog("exec file listxattr")
	resp.Xattr = listXattr(req.Size)
	return nil
}

// Readlink reads a symbolic link.
// This call is triggered when OS tries to see where symlink points
func (fi *File) Readlink(ctx context.Context, req *fuse.ReadlinkRequest) (string, error) {
	info, err := fi.m.fs.Stat(fi.path)
	if err != nil {
		return "/brig/backend/ipfs/", err
	}
	return fmt.Sprintf("/brig/backend/ipfs/%s", info.BackendHash), nil
}

// Compile time checks to see which interfaces we implement:
// Please update this list when modifying code here.
var _ = fs.Node(&File{})
var _ = fs.NodeFsyncer(&File{})
var _ = fs.NodeGetxattrer(&File{})
var _ = fs.NodeListxattrer(&File{})
var _ = fs.NodeOpener(&File{})
var _ = fs.NodeSetattrer(&File{})
var _ = fs.NodeReadlinker(&File{})

// Other interfaces are available, but currently not needed or make sense:
// var _ = fs.NodeRenamer(&File{})
// var _ = fs.NodeReadlinker(&File{})
// var _ = fs.NodeRemover(&File{})
// var _ = fs.NodeRemovexattrer(&File{})
// var _ = fs.NodeRequestLookuper(&File{})
// var _ = fs.NodeAccesser(&File{})
// var _ = fs.NodeForgetter(&File{})
// var _ = fs.NodeGetattrer(&File{})
// var _ = fs.NodeLinker(&File{})
// var _ = fs.NodeMkdirer(&File{})
// var _ = fs.NodeMknoder(&File{})
// var _ = fs.NodeSetxattrer(&File{})
// var _ = fs.NodeStringLookuper(&File{})
// var _ = fs.NodeSymlinker(&File{})
