package installer

import (
	"path/filepath"

	"github.com/fishworks/fish"
	"github.com/fishworks/fish/pkg/osutil"
	"github.com/fishworks/fish/pkg/rig"
)

// LocalInstaller installs rigs from the filesystem
type LocalInstaller struct {
	Source string
	Home   fish.Home
}

// NewLocalInstaller creates a new LocalInstaller
func NewLocalInstaller(source string, home fish.Home) (*LocalInstaller, error) {
	i := &LocalInstaller{
		Source: source,
		Home:   home,
	}

	return i, nil
}

// Install creates a symlink to the rig directory
func (i *LocalInstaller) Install() error {
	if !isRig(i.Source) {
		return rig.ErrMissingMetadata
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
	return filepath.Join(i.Home.Rigs(), filepath.Base(i.Source))
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
	return osutil.SymlinkWithFallback(origin, dest)
}
