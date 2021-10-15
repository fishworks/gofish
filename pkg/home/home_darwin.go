//go:build darwin
// +build darwin

package home

// HomePrefix is the base path to Fish's configuration directory.
const HomePrefix = "/usr/local"

// Cache returns the path to the cache.
func Cache() string {
	return userpath.Path("Library", "Caches", "gofish")
}
