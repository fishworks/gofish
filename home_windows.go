// +build windows

package gofish

import (
	"os"
)

// HomePrefix is the base path to Fish's configuration directory.
const HomePrefix = "C:\\ProgramData"

// BinPath is the path to where executables should be installed by gofish.
const BinPath = HomePrefix + "\\bin"

// UserHomePath is the path to $HOME
var UserHomePath = homedir()

func homedir() string {
	if home := os.Getenv("HOME"); len(home) > 0 {
		return home
	}
	return os.Getenv("USERPROFILE")
}

// Cache returns the path to the cache.
func (h UserHome) Cache() string {
	return h.Path("AppData", "Local", "Fish")
}
