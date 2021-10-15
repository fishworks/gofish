//go:build windows
// +build windows

package home

// HomePrefix is the base path to Fish's configuration directory.
const HomePrefix = "C:\\ProgramData"

func Cache() string {
	return userpath.Path("AppData", "Local", "gofish")
}
