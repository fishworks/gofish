//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

func ensureDirectories(dirs []string) error {
	var isRoot bool = false

	curUser, err := user.Current()
	if err != nil {
		return err
	}

	i, err := strconv.Atoi(curUser.Uid)
	if err != nil {
		return err
	} else if i == 0 {
		isRoot = true
	}

	curGroup, err := user.LookupGroupId(curUser.Gid)
	if err != nil {
		return err
	}

	fmt.Printf("The following new directories will be created:\n")
	fmt.Println(strings.Join(dirs, "\n"))
	var cmd *exec.Cmd
	for _, dir := range dirs {
		if fi, err := os.Stat(dir); err != nil {
			if isRoot {
				cmd = exec.Command("mkdir", dir)
			} else {
				cmd = exec.Command("sudo", "mkdir", dir)
			}
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			}
		} else if !fi.IsDir() {
			return fmt.Errorf("%s must be a directory", dir)
		}
		if isRoot {
			cmd = exec.Command("chown", fmt.Sprintf("%s:%s", curUser.Username, curGroup.Name), dir)
		} else {
			cmd = exec.Command("sudo", "chown", fmt.Sprintf("%s:%s", curUser.Username, curGroup.Name), dir)
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
