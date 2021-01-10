package repo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	e "github.com/pkg/errors"
	"github.com/sahib/brig/defaults"
	"github.com/sahib/config"
	log "github.com/sirupsen/logrus"
)

func touch(path string) error {
	fd, err := os.OpenFile(path, os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	return fd.Close()
}

// InitOptions sum up the option that we can pass to Init()
type InitOptions struct {
	// BaseFolder is where the repository is located.
	BaseFolder string
	// Owner is the owner id of the repository.
	Owner string
	// Password is the password used to lock the repository.
	Password string
	// BackendName says what backend we should use.
	BackendName string
	// DaemonURL is the URL that will be used for the brig daemon.
	DaemonURL string
}

// IsValidBackendName tells you if `name` is a valid backend name.
func IsValidBackendName(name string) bool {
	switch name {
	case "mock", "httpipfs":
		return true
	default:
		return false
	}
}

// Validate checks if the options are valid.
func (opts InitOptions) Validate() error {
	if !IsValidBackendName(opts.BackendName) {
		return fmt.Errorf("invalid backend name: %v", opts.BackendName)
	}

	if len(opts.Owner) == 0 {
		return fmt.Errorf("owner may not be empty")
	}

	return nil
}

// Init will create a new repository on disk at `baseFolder`.
// `owner` will be the new owner and should be something like user@domain/resource.
// `backendName` is the name of the backend, either "ipfs" or "mock".
// `daemonPort` is the port of the local daemon.
func Init(opts InitOptions) error {
	if err := opts.Validate(); err != nil {
		return err
	}

	// The basefolder has to exist:
	info, err := os.Stat(opts.BaseFolder)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(opts.BaseFolder, 0700); err != nil {
			return err
		}
	} else if info.Mode().IsDir() {
		children, err := ioutil.ReadDir(opts.BaseFolder)
		if err != nil {
			return err
		}

		if len(children) > 0 {
			log.Warningf("`%s` is a directory and exists", opts.BaseFolder)
		}
	} else {
		return fmt.Errorf("`%s` is a file (should be a directory)", opts.BaseFolder)
	}

	// Create (empty) folders:
	for _, emptyFolder := range []string{"metadata", "data"} {
		absFolder := filepath.Join(opts.BaseFolder, emptyFolder)
		if err := os.Mkdir(absFolder, 0700); err != nil {
			return e.Wrapf(err, "Failed to create dir: %v (repo exists?)", absFolder)
		}
	}

	if err := touch(filepath.Join(opts.BaseFolder, "remotes.yml")); err != nil {
		return e.Wrapf(err, "Failed touch remotes.yml")
	}

	if err := touch(filepath.Join(opts.BaseFolder, "INIT_TAG")); err != nil {
		return e.Wrapf(err, "Failed touch INIT_TAG")
	}

	ownerPath := filepath.Join(opts.BaseFolder, "OWNER")
	if err := ioutil.WriteFile(
		ownerPath,
		[]byte(opts.Owner),
		0644,
	); err != nil {
		return err
	}

	backendNamePath := filepath.Join(opts.BaseFolder, "BACKEND")
	if err := ioutil.WriteFile(
		backendNamePath,
		[]byte(opts.BackendName),
		0644,
	); err != nil {
		return err
	}

	// For future use: If we ever need to migrate the repo.
	versionPath := filepath.Join(opts.BaseFolder, "VERSION")
	if err := ioutil.WriteFile(
		versionPath,
		[]byte("1"),
		0644,
	); err != nil {
		return err
	}

	// Create a default config, only with the default keys applied:
	cfg, err := config.Open(nil, defaults.Defaults, config.StrictnessPanic)
	if err != nil {
		return err
	}

	if err := cfg.SetString("daemon.url", opts.DaemonURL); err != nil {
		return err
	}

	configPath := filepath.Join(opts.BaseFolder, "config.yml")
	if err := config.ToYamlFile(configPath, cfg); err != nil {
		return e.Wrap(err, "Failed to setup default config")
	}

	dataFolder := filepath.Join(opts.BaseFolder, "data", opts.BackendName)
	if err := os.MkdirAll(dataFolder, 0700); err != nil {
		return e.Wrap(err, "Failed to setup dirs for backend")
	}

	// Create initial key pair:
	if err := createKeyPair(opts.Owner, opts.BaseFolder, 2048); err != nil {
		return e.Wrap(err, "Failed to setup gpg keys")
	}

	passwdFile := filepath.Join(opts.BaseFolder, "passwd")
	if err := ioutil.WriteFile(
		passwdFile,
		[]byte(opts.Owner),
		0644,
	); err != nil {
		return err
	}

	// passwd is used to verify the user password,
	// so it needs to be locked only once on init and
	// kept out otherwise from the locking machinery.
	return lockFile(
		passwdFile,
		keyFromPassword(opts.Owner, opts.Password),
	)
}

// OverwriteConfigKey allows to overwrite a single key/val pair in the config,
// without requiring a running daemon or an opened repository.
// It is not performant and should be use with care.
func OverwriteConfigKey(repoPath string, key string, val interface{}) error {
	configPath := filepath.Join(repoPath, "config.yml")
	cfg, err := defaults.OpenMigratedConfig(configPath)
	if err != nil {
		return e.Wrapf(err, "failed to set ipfs port")
	}

	if err := cfg.Set(key, val); err != nil {
		return err
	}

	fd, err := os.OpenFile(configPath, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	defer fd.Close()

	return cfg.Save(config.NewYamlEncoder(fd))
}
