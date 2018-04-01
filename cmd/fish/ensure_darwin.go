// +build darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// ensureDirectories on MacOS differs from UNIX in the sense that it chowns the directories as the
// admin user for Homebrew compatibility.
func ensureDirectories(dirs []string) error {
	curUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("Could not determine current user: %s", err)
	}
	fmt.Printf("The following new directories will be created and will have their owner set to %s:\n", curUser.Name)
	fmt.Println(strings.Join(dirs, "\n"))
	for _, dir := range dirs {
		if fi, err := os.Stat(dir); err != nil {
			cmd := exec.Command("sudo", "mkdir", "-p", dir)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			}
		} else if !fi.IsDir() {
			return fmt.Errorf("%s must be a directory", dir)
		}
		cmd := exec.Command("sudo", "chown", "-R", fmt.Sprintf("%s:%s", curUser.Name, "admin"), dir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
