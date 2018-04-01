package installer

import (
	"os"
	"path"
	"path/filepath"

	"github.com/fishworks/fish"
	"github.com/fishworks/fish/pkg/rig"
)

// Installer provides an interface for installing client rigs.
type Installer interface {
	// Install adds a rig to a path
	Install() error
	// Path is the directory of the installed rig.
	Path() string
	// Update updates a rig.
	Update() error
}

// Install installs a rig.
func Install(i Installer) error {
	basePath := path.Dir(i.Path())
	if _, pathErr := os.Stat(basePath); os.IsNotExist(pathErr) {
		if err := os.MkdirAll(basePath, 0755); err != nil {
			return err
		}
	}

	if _, pathErr := os.Stat(i.Path()); !os.IsNotExist(pathErr) {
		return i.Update()
	}

	return i.Install()
}

// Update updates a rig.
func Update(i Installer) error {
	if _, pathErr := os.Stat(i.Path()); os.IsNotExist(pathErr) {
		return rig.ErrDoesNotExist
	}

	return i.Update()
}

// FindSource determines the correct Installer for the given source.
func FindSource(location string, home fish.Home) (Installer, error) {
	installer, err := existingVCSRepo(location, home)
	if err != nil && err.Error() == "Cannot detect VCS" {
		return installer, rig.ErrMissingSource
	}
	return installer, err
}

// New determines and returns the correct Installer for the given source
func New(source, version string, home fish.Home) (Installer, error) {
	if isLocalReference(source) {
		return NewLocalInstaller(source, home)
	}

	return NewVCSInstaller(source, version, home)
}

// isLocalReference checks if the source exists on the filesystem.
func isLocalReference(source string) bool {
	_, err := os.Stat(source)
	return err == nil
}

// isRig checks if the directory contains a "Food" directory.
func isRig(dirname string) bool {
	_, err := os.Stat(filepath.Join(dirname, "Food"))
	return err == nil
}
