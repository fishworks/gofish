package installer

import (
	"os"
	"path/filepath"

	"github.com/fishworks/gofish"
	"github.com/fishworks/gofish/pkg/rig"
)

// LocalInstaller installs rigs from the filesystem
type LocalInstaller struct {
	Source string
	Home   gofish.Home
	Name   string
}

// NewLocalInstaller creates a new LocalInstaller
func NewLocalInstaller(source string, name string, home gofish.Home) (*LocalInstaller, error) {
	i := &LocalInstaller{
		Source: source,
		Home:   home,
		Name:   name,
	}

	if i.Name == "" {
		i.Name = filepath.Base(i.Source)
	}

	return i, nil
}

// Install creates a symlink to the rig directory
func (i *LocalInstaller) Install() error {
	if !isRig(i.Source) {
		return rig.ErrDoesNotExist
	}

	src, err := filepath.Abs(i.Source)
	if err != nil {
		return err
	}

	return i.link(src)
}

// Path is where the rig will be installed into.
func (i *LocalInstaller) Path() string {
	if i.Source == "" {
		return ""
	}
	return filepath.Join(i.Home.Rigs(), i.Name)
}

// Update updates a local repository, which is a no-op.
func (i *LocalInstaller) Update() error {
	return nil
}

// link creates a symlink from the rig source to the rig directory
func (i *LocalInstaller) link(from string) error {
	origin, err := filepath.Abs(from)
	if err != nil {
		return err
	}
	dest, err := filepath.Abs(i.Path())
	if err != nil {
		return err
	}
	return os.Symlink(origin, dest)
}
