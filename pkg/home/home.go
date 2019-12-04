package home

import (
	"os"
	"path"
	"path/filepath"

	"github.com/fishworks/gofish/pkg/lazypath"
)

// DefaultHomeEnvVar defines the environment variable used to look up the home directory.
const DefaultHomeEnvVar = "GOFISH_HOME"

var (
	homepath = lazypath.LazyPath{
		EnvironmentVariable: DefaultHomeEnvVar,
		DefaultFn:           defaultHome,
	}

	userpath = lazypath.LazyPath{
		EnvironmentVariable: "HOME",
		DefaultFn:           defaultUserHome,
	}

	binpath = lazypath.LazyPath{
		EnvironmentVariable: "GOFISH_BINPATH",
		DefaultFn:           defaultBinPath,
	}
)

func defaultHome() string {
	return filepath.Join(HomePrefix, "gofish")
}

func defaultUserHome() string {
	if home := os.Getenv("HOME"); len(home) > 0 {
		return home
	}
	return os.Getenv("USERPROFILE")
}

func defaultBinPath() string {
	return filepath.Join(HomePrefix, "bin")
}

// Barrel returns the path to the fish barrel.
func Barrel() string {
	return homepath.Path("Barrel")
}

// Rigs returns the path to the fishing rigs.
func Rigs() string {
	return homepath.Path("Rigs")
}

// String returns Home as a string.
//
// Implements fmt.Stringer.
func String() string {
	return homepath.Path("")
}

// UserHome returns the home path.
func UserHome() string {
	return userpath.Path("")
}

// BinPath is the path where executables should be installed by gofish.
func BinPath() string {
	return binpath.Path("")
}

// DefaultRig returns the name of the default fishing rig.
func DefaultRig() string {
	return path.Join("github.com", "fishworks", "fish-food")
}
