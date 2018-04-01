// +build !windows

package fish

import (
	"os"
)

const HomePrefix = "/usr/local"

const BinPath = HomePrefix + "/bin"

var UserHomePath = os.Getenv("HOME")

// Cache returns the path to the cache.
func (h UserHome) Cache() string {
	return h.Path("Library", "Caches", "Fish")
}
