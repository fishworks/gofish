// +build !windows,!darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ensureDirectories(dirs []string) error {
	curUser := os.Getenv("USER")
	if curUser == "" {
		return fmt.Errorf("Could not determine current user: $USER is not present in the environment")
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
		cmd := exec.Command("sudo", "chown", fmt.Sprintf("%s:%s", curUser, curUser), dir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
