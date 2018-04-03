// +build !windows,!darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func ensureDirectories(dirs []string) error {
	curUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("Could not determine current user: %s", err)
	}
	fmt.Printf("The following new directories will be created:\n")
	fmt.Println(strings.Join(dirs, "\n"))
	for _, dir := range dirs {
		if fi, err := os.Stat(dir); err != nil {
			cmd := exec.Command("sudo", "mkdir", dir)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			}
		} else if !fi.IsDir() {
			return fmt.Errorf("%s must be a directory", dir)
		}
		cmd := exec.Command("sudo", "chown", fmt.Sprintf("%s:%s", curUser.Uid, curUser.Gid), dir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
