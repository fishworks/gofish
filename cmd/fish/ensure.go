package main

import (
	"time"

	"github.com/fishworks/fish/pkg/ohai"
)

// ensureFood checks to see if the default fish food exists.
//
// If the pack does not exist, this function will create it.
// If it does, it will update to the latest.
func ensureFood() error {
	ohai.Ohailn("Installing default fish food...")

	addArgs := []string{
		"add",
		"github.com/fishworks/fish-food",
	}

	rigCmd, _, err := rootCmd.Find([]string{"rig"})
	if err != nil {
		return err
	}

	start := time.Now()
	if err := rigCmd.RunE(rigCmd, addArgs); err != nil {
		return err
	}
	t := time.Now()
	ohai.Successf("fishworks/fish-food: installed in %s\n", t.Sub(start).String())
	return nil
}
