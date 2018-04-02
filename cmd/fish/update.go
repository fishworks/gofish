package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/fishworks/fish"
	"github.com/fishworks/fish/pkg/ohai"
	"github.com/fishworks/fish/pkg/rig/installer"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

type updateCmd struct{}

func newUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update rigs",
		RunE: func(_ *cobra.Command, _ []string) error {
			return updateRigs()
		},
	}
	return cmd
}

func updateRigs() error {
	start := time.Now()
	home := fish.Home(fish.HomePath)
	rigs := findRigs(home.Rigs())
	for _, rig := range rigs {
		i, err := installer.FindSource(filepath.Join(home.Rigs(), rig), home)
		if err != nil {
			return err
		}
		if err := installer.Update(i); err != nil {
			return err
		}
	}
	t := time.Now()
	ohai.Ohailn("Rigs updated!")
	table := uitable.New()
	table.AddRow("NAME")
	for _, rig := range rigs {
		table.AddRow(rig)
	}
	fmt.Printf("%s\n\n", table)
	ohai.Successf("rigs updated in %s\n", t.Sub(start).String())
	return nil
}
