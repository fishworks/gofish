//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ensureDirectories(dirs []string) error {
	userCmd := exec.Command("id", "-un")
	userOutput, err := userCmd.Output()
	if err != nil {
		return err
	}
	// strip the newline character from the end
	curUser := strings.TrimSuffix(string(userOutput), "\n")

	groupCmd := exec.Command("id", "-gn", curUser)
	groupOutput, err := groupCmd.Output()
	if err != nil {
		return err
	}
	// strip the newline character from the end
	curGroup := strings.TrimSuffix(string(groupOutput), "\n")

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
		cmd := exec.Command("sudo", "chown", fmt.Sprintf("%s:%s", curUser, curGroup), dir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
