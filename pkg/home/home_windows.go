//go:build windows
// +build windows

package home

// HomePrefix is the base path to Fish's configuration directory.
const HomePrefix = "C:\\ProgramData"

func Cache() string {
	return userpath.Path("AppData", "Local", "gofish")
}

// GPGNetrc returns the path to an encrypted netrc file.
func GPGNetrc() string {
	return userpath.Path("_netrc.gpg")
}

// Netrc returns the path to a netrc file.
func Netrc() string {
	return userpath.Path("_netrc")
}
