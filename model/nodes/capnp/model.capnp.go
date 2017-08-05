// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Person might be any brig user
type Person struct{ capnp.Struct }

// Person_TypeID is the unique identifier for the type Person.
const Person_TypeID = 0xf736dd278ea58545

func NewPerson(s *capnp.Segment) (Person, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Person{st}, err
}

func NewRootPerson(s *capnp.Segment) (Person, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Person{st}, err
}

func ReadRootPerson(msg *capnp.Message) (Person, error) {
	root, err := msg.RootPtr()
	return Person{root.Struct()}, err
}

func (s Person) String() string {
	str, _ := text.Marshal(0xf736dd278ea58545, s.Struct)
	return str
}

func (s Person) Ident() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Person) HasIdent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Person) IdentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Person) SetIdent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Person) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Person) HasHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Person) SetHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

// Person_List is a list of Person.
type Person_List struct{ capnp.List }

// NewPerson creates a new list of Person.
func NewPerson_List(s *capnp.Segment, sz int32) (Person_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Person_List{l}, err
}

func (s Person_List) At(i int) Person { return Person{s.List.Struct(i)} }

func (s Person_List) Set(i int, v Person) error { return s.List.SetStruct(i, v.Struct) }

func (s Person_List) String() string {
	str, _ := text.MarshalList(0xf736dd278ea58545, s.List)
	return str
}

// Person_Promise is a wrapper for a Person promised by a client call.
type Person_Promise struct{ *capnp.Pipeline }

func (p Person_Promise) Struct() (Person, error) {
	s, err := p.Pipeline.Struct()
	return Person{s}, err
}

// Commit is a set of changes to nodes
type Commit struct{ capnp.Struct }
type Commit_merge Commit

// Commit_TypeID is the unique identifier for the type Commit.
const Commit_TypeID = 0x8da013c66e545daf

func NewCommit(s *capnp.Segment) (Commit, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return Commit{st}, err
}

func NewRootCommit(s *capnp.Segment) (Commit, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return Commit{st}, err
}

func ReadRootCommit(msg *capnp.Message) (Commit, error) {
	root, err := msg.RootPtr()
	return Commit{root.Struct()}, err
}

func (s Commit) String() string {
	str, _ := text.Marshal(0x8da013c66e545daf, s.Struct)
	return str
}

func (s Commit) Message() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Commit) HasMessage() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Commit) MessageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Commit) SetMessage(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Commit) Parent() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Commit) HasParent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Commit) SetParent(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s Commit) Root() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s Commit) HasRoot() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Commit) SetRoot(v []byte) error {
	return s.Struct.SetData(2, v)
}

func (s Commit) Merge() Commit_merge { return Commit_merge(s) }

func (s Commit_merge) IsMerge() bool {
	return s.Struct.Bit(0)
}

func (s Commit_merge) SetIsMerge(v bool) {
	s.Struct.SetBit(0, v)
}

func (s Commit_merge) With() (Person, error) {
	p, err := s.Struct.Ptr(3)
	return Person{Struct: p.Struct()}, err
}

func (s Commit_merge) HasWith() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Commit_merge) SetWith(v Person) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewWith sets the with field to a newly
// allocated Person struct, preferring placement in s's segment.
func (s Commit_merge) NewWith() (Person, error) {
	ss, err := NewPerson(s.Struct.Segment())
	if err != nil {
		return Person{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Commit_merge) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return []byte(p.Data()), err
}

func (s Commit_merge) HasHash() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Commit_merge) SetHash(v []byte) error {
	return s.Struct.SetData(4, v)
}

// Commit_List is a list of Commit.
type Commit_List struct{ capnp.List }

// NewCommit creates a new list of Commit.
func NewCommit_List(s *capnp.Segment, sz int32) (Commit_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	return Commit_List{l}, err
}

func (s Commit_List) At(i int) Commit { return Commit{s.List.Struct(i)} }

func (s Commit_List) Set(i int, v Commit) error { return s.List.SetStruct(i, v.Struct) }

func (s Commit_List) String() string {
	str, _ := text.MarshalList(0x8da013c66e545daf, s.List)
	return str
}

// Commit_Promise is a wrapper for a Commit promised by a client call.
type Commit_Promise struct{ *capnp.Pipeline }

func (p Commit_Promise) Struct() (Commit, error) {
	s, err := p.Pipeline.Struct()
	return Commit{s}, err
}

func (p Commit_Promise) Merge() Commit_merge_Promise { return Commit_merge_Promise{p.Pipeline} }

// Commit_merge_Promise is a wrapper for a Commit_merge promised by a client call.
type Commit_merge_Promise struct{ *capnp.Pipeline }

func (p Commit_merge_Promise) Struct() (Commit_merge, error) {
	s, err := p.Pipeline.Struct()
	return Commit_merge{s}, err
}

func (p Commit_merge_Promise) With() Person_Promise {
	return Person_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

// A single directory entry
type DirEntry struct{ capnp.Struct }

// DirEntry_TypeID is the unique identifier for the type DirEntry.
const DirEntry_TypeID = 0x8b15ee76774b1f9d

func NewDirEntry(s *capnp.Segment) (DirEntry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DirEntry{st}, err
}

func NewRootDirEntry(s *capnp.Segment) (DirEntry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DirEntry{st}, err
}

func ReadRootDirEntry(msg *capnp.Message) (DirEntry, error) {
	root, err := msg.RootPtr()
	return DirEntry{root.Struct()}, err
}

func (s DirEntry) String() string {
	str, _ := text.Marshal(0x8b15ee76774b1f9d, s.Struct)
	return str
}

func (s DirEntry) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s DirEntry) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DirEntry) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s DirEntry) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s DirEntry) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s DirEntry) HasHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s DirEntry) SetHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

// DirEntry_List is a list of DirEntry.
type DirEntry_List struct{ capnp.List }

// NewDirEntry creates a new list of DirEntry.
func NewDirEntry_List(s *capnp.Segment, sz int32) (DirEntry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return DirEntry_List{l}, err
}

func (s DirEntry_List) At(i int) DirEntry { return DirEntry{s.List.Struct(i)} }

func (s DirEntry_List) Set(i int, v DirEntry) error { return s.List.SetStruct(i, v.Struct) }

func (s DirEntry_List) String() string {
	str, _ := text.MarshalList(0x8b15ee76774b1f9d, s.List)
	return str
}

// DirEntry_Promise is a wrapper for a DirEntry promised by a client call.
type DirEntry_Promise struct{ *capnp.Pipeline }

func (p DirEntry_Promise) Struct() (DirEntry, error) {
	s, err := p.Pipeline.Struct()
	return DirEntry{s}, err
}

// Directory contains one or more directories or files
type Directory struct{ capnp.Struct }

// Directory_TypeID is the unique identifier for the type Directory.
const Directory_TypeID = 0xe24c59306c829c01

func NewDirectory(s *capnp.Segment) (Directory, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Directory{st}, err
}

func NewRootDirectory(s *capnp.Segment) (Directory, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Directory{st}, err
}

func ReadRootDirectory(msg *capnp.Message) (Directory, error) {
	root, err := msg.RootPtr()
	return Directory{root.Struct()}, err
}

func (s Directory) String() string {
	str, _ := text.Marshal(0xe24c59306c829c01, s.Struct)
	return str
}

func (s Directory) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Directory) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Directory) Parent() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Directory) HasParent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Directory) ParentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Directory) SetParent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Directory) Children() (DirEntry_List, error) {
	p, err := s.Struct.Ptr(1)
	return DirEntry_List{List: p.List()}, err
}

func (s Directory) HasChildren() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Directory) SetChildren(v DirEntry_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewChildren sets the children field to a newly
// allocated DirEntry_List, preferring placement in s's segment.
func (s Directory) NewChildren(n int32) (DirEntry_List, error) {
	l, err := NewDirEntry_List(s.Struct.Segment(), n)
	if err != nil {
		return DirEntry_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Directory_List is a list of Directory.
type Directory_List struct{ capnp.List }

// NewDirectory creates a new list of Directory.
func NewDirectory_List(s *capnp.Segment, sz int32) (Directory_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Directory_List{l}, err
}

func (s Directory_List) At(i int) Directory { return Directory{s.List.Struct(i)} }

func (s Directory_List) Set(i int, v Directory) error { return s.List.SetStruct(i, v.Struct) }

func (s Directory_List) String() string {
	str, _ := text.MarshalList(0xe24c59306c829c01, s.List)
	return str
}

// Directory_Promise is a wrapper for a Directory promised by a client call.
type Directory_Promise struct{ *capnp.Pipeline }

func (p Directory_Promise) Struct() (Directory, error) {
	s, err := p.Pipeline.Struct()
	return Directory{s}, err
}

type File struct{ capnp.Struct }

// File_TypeID is the unique identifier for the type File.
const File_TypeID = 0x8ea7393d37893155

func NewFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return File{st}, err
}

func NewRootFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return File{st}, err
}

func ReadRootFile(msg *capnp.Message) (File, error) {
	root, err := msg.RootPtr()
	return File{root.Struct()}, err
}

func (s File) String() string {
	str, _ := text.Marshal(0x8ea7393d37893155, s.Struct)
	return str
}

func (s File) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s File) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s File) Parent() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s File) HasParent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s File) ParentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s File) SetParent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s File) Key() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s File) HasKey() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s File) SetKey(v []byte) error {
	return s.Struct.SetData(1, v)
}

// File_List is a list of File.
type File_List struct{ capnp.List }

// NewFile creates a new list of File.
func NewFile_List(s *capnp.Segment, sz int32) (File_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return File_List{l}, err
}

func (s File_List) At(i int) File { return File{s.List.Struct(i)} }

func (s File_List) Set(i int, v File) error { return s.List.SetStruct(i, v.Struct) }

func (s File_List) String() string {
	str, _ := text.MarshalList(0x8ea7393d37893155, s.List)
	return str
}

// File_Promise is a wrapper for a File promised by a client call.
type File_Promise struct{ *capnp.Pipeline }

func (p File_Promise) Struct() (File, error) {
	s, err := p.Pipeline.Struct()
	return File{s}, err
}

// Ghost indicates that a certain node was at this path once
type Ghost struct{ capnp.Struct }

// Ghost_TypeID is the unique identifier for the type Ghost.
const Ghost_TypeID = 0x80c828d7e89c12ea

func NewGhost(s *capnp.Segment) (Ghost, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Ghost{st}, err
}

func NewRootGhost(s *capnp.Segment) (Ghost, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Ghost{st}, err
}

func ReadRootGhost(msg *capnp.Message) (Ghost, error) {
	root, err := msg.RootPtr()
	return Ghost{root.Struct()}, err
}

func (s Ghost) String() string {
	str, _ := text.Marshal(0x80c828d7e89c12ea, s.Struct)
	return str
}

func (s Ghost) NodeType() uint8 {
	return s.Struct.Uint8(0)
}

func (s Ghost) SetNodeType(v uint8) {
	s.Struct.SetUint8(0, v)
}

// Ghost_List is a list of Ghost.
type Ghost_List struct{ capnp.List }

// NewGhost creates a new list of Ghost.
func NewGhost_List(s *capnp.Segment, sz int32) (Ghost_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Ghost_List{l}, err
}

func (s Ghost_List) At(i int) Ghost { return Ghost{s.List.Struct(i)} }

func (s Ghost_List) Set(i int, v Ghost) error { return s.List.SetStruct(i, v.Struct) }

func (s Ghost_List) String() string {
	str, _ := text.MarshalList(0x80c828d7e89c12ea, s.List)
	return str
}

// Ghost_Promise is a wrapper for a Ghost promised by a client call.
type Ghost_Promise struct{ *capnp.Pipeline }

func (p Ghost_Promise) Struct() (Ghost, error) {
	s, err := p.Pipeline.Struct()
	return Ghost{s}, err
}

// Node is a node in the merkle dag of brig
type Node struct{ capnp.Struct }
type Node_Which uint16

const (
	Node_Which_commit    Node_Which = 0
	Node_Which_ghost     Node_Which = 1
	Node_Which_directory Node_Which = 2
	Node_Which_file      Node_Which = 3
)

func (w Node_Which) String() string {
	const s = "commitghostdirectoryfile"
	switch w {
	case Node_Which_commit:
		return s[0:6]
	case Node_Which_ghost:
		return s[6:11]
	case Node_Which_directory:
		return s[11:20]
	case Node_Which_file:
		return s[20:24]

	}
	return "Node_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Node_TypeID is the unique identifier for the type Node.
const Node_TypeID = 0xa629eb7f7066fae3

func NewNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return Node{st}, err
}

func NewRootNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return Node{st}, err
}

func ReadRootNode(msg *capnp.Message) (Node, error) {
	root, err := msg.RootPtr()
	return Node{root.Struct()}, err
}

func (s Node) String() string {
	str, _ := text.Marshal(0xa629eb7f7066fae3, s.Struct)
	return str
}

func (s Node) Which() Node_Which {
	return Node_Which(s.Struct.Uint16(0))
}
func (s Node) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Node) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Node) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Node) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Node) HasHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Node) SetHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s Node) ModTime() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Node) HasModTime() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Node) ModTimeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Node) SetModTime(v string) error {
	return s.Struct.SetText(2, v)
}

func (s Node) Commit() (Commit, error) {
	p, err := s.Struct.Ptr(3)
	return Commit{Struct: p.Struct()}, err
}

func (s Node) HasCommit() bool {
	if s.Struct.Uint16(0) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetCommit(v Commit) error {
	s.Struct.SetUint16(0, 0)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewCommit sets the commit field to a newly
// allocated Commit struct, preferring placement in s's segment.
func (s Node) NewCommit() (Commit, error) {
	s.Struct.SetUint16(0, 0)
	ss, err := NewCommit(s.Struct.Segment())
	if err != nil {
		return Commit{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Ghost() (Ghost, error) {
	p, err := s.Struct.Ptr(3)
	return Ghost{Struct: p.Struct()}, err
}

func (s Node) HasGhost() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetGhost(v Ghost) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewGhost sets the ghost field to a newly
// allocated Ghost struct, preferring placement in s's segment.
func (s Node) NewGhost() (Ghost, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewGhost(s.Struct.Segment())
	if err != nil {
		return Ghost{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Directory() (Directory, error) {
	p, err := s.Struct.Ptr(3)
	return Directory{Struct: p.Struct()}, err
}

func (s Node) HasDirectory() bool {
	if s.Struct.Uint16(0) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetDirectory(v Directory) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewDirectory sets the directory field to a newly
// allocated Directory struct, preferring placement in s's segment.
func (s Node) NewDirectory() (Directory, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := NewDirectory(s.Struct.Segment())
	if err != nil {
		return Directory{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) File() (File, error) {
	p, err := s.Struct.Ptr(3)
	return File{Struct: p.Struct()}, err
}

func (s Node) HasFile() bool {
	if s.Struct.Uint16(0) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetFile(v File) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewFile sets the file field to a newly
// allocated File struct, preferring placement in s's segment.
func (s Node) NewFile() (File, error) {
	s.Struct.SetUint16(0, 3)
	ss, err := NewFile(s.Struct.Segment())
	if err != nil {
		return File{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

// Node_List is a list of Node.
type Node_List struct{ capnp.List }

// NewNode creates a new list of Node.
func NewNode_List(s *capnp.Segment, sz int32) (Node_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return Node_List{l}, err
}

func (s Node_List) At(i int) Node { return Node{s.List.Struct(i)} }

func (s Node_List) Set(i int, v Node) error { return s.List.SetStruct(i, v.Struct) }

func (s Node_List) String() string {
	str, _ := text.MarshalList(0xa629eb7f7066fae3, s.List)
	return str
}

// Node_Promise is a wrapper for a Node promised by a client call.
type Node_Promise struct{ *capnp.Pipeline }

func (p Node_Promise) Struct() (Node, error) {
	s, err := p.Pipeline.Struct()
	return Node{s}, err
}

func (p Node_Promise) Commit() Commit_Promise {
	return Commit_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Node_Promise) Ghost() Ghost_Promise {
	return Ghost_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Node_Promise) Directory() Directory_Promise {
	return Directory_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Node_Promise) File() File_Promise {
	return File_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

const schema_9195d073cb5c5953 = "x\xda\x9c\x95_\x88T\xe5\x1b\xc7\x9f\xef\xfb\x9e\x99\xb3" +
	"\xfeX\x7f;\xc7w\x85\x02\x97y\x15#]ZS\xd3" +
	"JAVm\xa7\xd24\xf6\xad\xedB\xd0\xe88\xf3\xee" +
	"\x9c\x17g\xce\x19\xce9\xb5L\x10[A\x17\x15\x1a\x16" +
	"]\x08J\x7f\xa0\xba\xb4\x8b\x02\x13\x05\x8b\x0c3/*" +
	"\xfaG\xe0\x8d]\x14\x05]t\x13Hu\xe2\x9d\x99\x9d" +
	"\x99lV\xa1\xcb\xf3\x9c\xe7<\xf3y\x9e\xe7\xfb|g" +
	"\xfdQ\xb6\x9dm\xc8\xcd;DjS.\x9f\xfd\xbc\xec" +
	"\xf8O\xdf\xae\xb9\xf04\xa9e@\xf6\xf0\xbe\xfd\x9f%" +
	"\x9f\xbfz\x94\x1c\x97H\xdc\xc2>\x12\x13\xcc\x15\x13\xac" +
	"(\x0c\x9b$d'\x8a\x0f\xcc=\xf1\xeb\xf2\x17\xc9[" +
	"\xd6\x97\x9cc6\xfb0\xfb^\x1cc\xae8\xc6\x8a\xe2" +
	"<\x9b#d'\x0f\xcc\x84\x9f\x88\xd7\x0f_S;\x97" +
	"\xb3\xe9+\xf9E1\xc1]1\xc1\x8b\xe2\x00\xff\x91\x90" +
	"=\xb2\xe1\xf9\xbb\xb6my\xe7\xc8\xb5\xe9\xad\xea\xcb\x9d" +
	"\xd3b\xccq\xc5\x98S\x14\xdb\x9c\x93\x84\xec\x87\xab\xb3" +
	"\x8d\xf9_\xd6\xbe}Mz\xc9q\x1d\"\xf1\xb5sZ" +
	"\\v\\q\xd9)\xde\xe1\xe5^\x02!{f\xf3\xca" +
	"\xf0\xe5\xd1\xf1S\xa4nF\x1f\xdc\xf2\x9c\x0b\"q\"" +
	"\x7f\x95 \xde\xc8\xdb\xd28\xfelm\xfd\xbe=W\x06" +
	"\x92lq\xaf\x88\x92\xeb\x8a\x92[\x14O\xb96\xbd\xf4" +
	"\xdc[Gn\xbd|\xe7\xef\x83\xa6\xb2v\xe8\xa2\xd8<" +
	"\xe4\x8a\xcdCE\xe1\x0f\xcd\xd1\xddY\xd9o\x84\x8d\xdb" +
	"\xeb\x11\xab\xe8\xda\xba\xd6\xc3\xd6\xfb\x82(Ii\x1aP" +
	"\x0eX\xf6\xe8+\xaf\xa9\xb3\xdf\xbcp\x9e\x94\xc3\xb0\xe3" +
	"6`\x98h\x03\xbeD\xd6J\x93&\xccWL\xd9O" +
	"u\"\xd3\xc0O\xa5/\xcb:N}\x13\xca0\xaah" +
	"9\xe7'\xd2Oe\x1a\x98D6\xfc4\x90QX\x86" +
	"&R\x0ew\x88\x1c\x10yKw\x13\xa9a\x0eu\x13" +
	"Cf?\x9ai64\x11!O\x0cy\xc2 \xc4)" +
	"\x13\x97\xc2\x94\xc7\xcd\xc1\x94+Z\x94\x1e.f;d" +
	"b\xc2jM3Y1\xb1.\xa7Q\xdc\x94:L\xe3" +
	"&A\x0du\x11\xd6\x8e\x13\xa9\xd5\x1cj=\x83\x07\x8c" +
	"\xc2\x06'lp\x0d\x87\xda\xc40\x12\xfau\x8dab" +
	"\x18&\x8c\x04~\x12`)1,\x1dLwOT\xaf" +
	"\x9b\x94\x16\x19\xa1\xec\x8cp\x15\xb2v\xa24<\x91\xbe" +
	"Lt*\xa3YY\x0e\xfc\xb0j\xa7\x19\xc90r+" +
	":!R\x85.\xa9\xbf\x93H\xed\xe7PA\x1f\xa9\xde" +
	"J\xa4\x1e\xe3P5\x06\x8f\xb1Q0\"\xcfX\xfc\x0a" +
	"\x87j0\x80\xa3Op^}#\xb1\xf9\xbaN\x12\xbf" +
	"\xdamj\xb2\xe1\xc7:L\x17\xda\x1a\x89\xa3\xa8\xfbP" +
	"\xac\xeb\xb8\xaa\xbb\x9db\xa1\xd3\xc9\xc6\xd6{MM\x0f" +
	"n\xb3\xd0\xd9\xc1\xff\x08j\xb8\xcb_\xb2T\xdb9\xd4" +
	"\x1e\x86\x05\xfc]\x16\x7f\x8aCM[|\xb4\xf1\xf7\xae" +
	"\"R\xf7s\xa8\x19\x86\x91\xc4<\xa9\xb1\x84\x18\x96\xf4" +
	"@;\xdc\xee!\xdd\xfc\xd7.\xfa\x09\x1f\x8c*\x8b\x10" +
	"\xae\xee,b72\x9b$M\"\x1d\xbf-[\x13\xca" +
	"4\xd0\xb2\xae\xe3C5-+~\xd5n\xe6`l\xaa" +
	"\x04\xb5\xa2\xdb\xcc\xfb\xb6\x99w9\xd4\x99\xbee|`" +
	"\x83\xefq\xa8s}\xcb8k\xd7v\x8aC}\xcc0" +
	"\xc6\xb3\x8c\x8f\x82\x13y\x1f\xda\xce\xcfp\xa8\x0b\x0cc" +
	"\xce_6\xec\x10y\xe77\x12\xa9s\x1c\xea\x12\xc3X" +
	"\xeeO\x1b\xce\x11y\x9f>D\xa4.p\xa8\xaf\x18\xc6" +
	"\xf2\x7f\xd8p\x9e\xc8\xfb\xc2\xfe\xe0%\x0e\xf5\xdd\xf5t" +
	":_\x8f*3\xa6\xf7r\xb2\xdc\xd2\x1e\x0a=\xeb!" +
	"\xa0@(V\xed]\xa3\xd0\xf3\xe2v<[8 B" +
	"\x13\x85\x9e)\xb5\xdf\x8e\xcc\x9a\x9aF\xa1\xe7\x9a\x9d\x8f" +
	"\x16\xbd\x8e\xc9u-Q\xb5\xa5\xd1\x9eFigO\x1b" +
	"\x1e:\xb3\xd85\xde/\x0e\xa7=\x89\xbd\xe3=q\xcc" +
	"\x9bd\xaf\xad\x04\x10\x03\x08#s&\x0dP\xe8\xd9`" +
	"\x07\xf0FG;\xd5j\xcf\x8d\x16\xf3\x945\x1d\xb5\xbc" +
	"\x89l\xaa3\x89\\S\x96\xa3\xd0\x9a]\"\xa3P\xcb" +
	"(\x96\xf5(\xd6]\xab1:\xb1\xb1Y\xe3\xd6Zg" +
	"\xfc_\xce\xc0\x9a\xe3\x9e\xf6\xbd_\xef\x0c\xb2r`j" +
	"\x95X\x87\xd68\xffO\x98\xe6@\xa1\xf7\x07I\xb0\xc1" +
	"AmO\xeb8\x89\xc2\xc5\xbcj\xc1H\x7f\xcb\xday" +
	"\xb2\xceL5H\xe5A-\xfd\xb0\xd9\xba\x88\xa2|<" +
	"\xd11Q\xbf\x9dn\xbc\x81\x9d\x16M\xa5\x8f\xfd\x1f\xab" +
	"\xf9;\x00\x00\xff\xff\xe2\xc6\xfb\x9a"

func init() {
	schemas.Register(schema_9195d073cb5c5953,
		0x80c828d7e89c12ea,
		0x8b15ee76774b1f9d,
		0x8da013c66e545daf,
		0x8ea7393d37893155,
		0xa629eb7f7066fae3,
		0xb82a14926e213581,
		0xe24c59306c829c01,
		0xf736dd278ea58545)
}
