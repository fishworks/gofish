//go:build windows
// +build windows

package main

import (
	"fmt"
	"os"
	"strings"
)

func ensureDirectories(dirs []string) error {
	fmt.Println("The following new directories will be created:")
	fmt.Println(strings.Join(dirs, "\n"))
	for _, dir := range dirs {
		if fi, err := os.Stat(dir); err != nil {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("Could not create %s: %s", dir, err)
			}
		} else if !fi.IsDir() {
			return fmt.Errorf("%s must be a directory", dir)
		}
	}
	return nil
}
