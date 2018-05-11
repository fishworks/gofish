package gofish

import (
	"path"
	"path/filepath"
)

type (
	// Home designates where fish should store data.
	Home string
	// UserHome designates the current user's home directory.
	UserHome string
)

// HomePath is the path to Fish's configuration directory.
var HomePath = filepath.Join(HomePrefix, "Fish")

// Path returns Home with elements appended.
func (h Home) Path(elem ...string) string {
	p := []string{h.String()}
	p = append(p, elem...)
	return filepath.Join(p...)
}

// Barrel returns the path to the fish barrel.
func (h Home) Barrel() string {
	return h.Path("Barrel")
}

// Rigs returns the path to the fishing rigs.
func (h Home) Rigs() string {
	return h.Path("Rigs")
}

// DefaultRig returns the name of the default fishing rig.
func (h Home) DefaultRig() string {
	return path.Join("github.com", "fishworks", "fish-food")
}

// String returns Home as a string.
//
// Implements fmt.Stringer.
func (h Home) String() string {
	return string(h)
}

// Path returns Home with elements appended.
func (h UserHome) Path(elem ...string) string {
	p := []string{h.String()}
	p = append(p, elem...)
	return filepath.Join(p...)
}

// String returns Home as a string.
//
// Implements fmt.Stringer.
func (h UserHome) String() string {
	return string(h)
}
