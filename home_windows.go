// +build windows

package fish

import (
	"os"
)

const HomePrefix = "C:\\"

const BinPath = HomePrefix

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
