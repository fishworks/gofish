//go:build darwin
// +build darwin

package home

// HomePrefix is the base path to Fish's configuration directory.
const HomePrefix = "/usr/local"

// Cache returns the path to the cache.
func Cache() string {
	return userpath.Path("Library", "Caches", "gofish")
}

// GPGNetrc returns the path to an encrypted netrc file.
func GPGNetrc() string {
	return userpath.Path(".netrc.gpg")
}

// Netrc returns the path to a netrc file.
func Netrc() string {
	return userpath.Path(".netrc")
}
