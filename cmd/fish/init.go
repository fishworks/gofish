package main

import (
	"path/filepath"

	"github.com/fishworks/fish"
	"github.com/spf13/cobra"
)

const (
	initDesc = `
Initializes fish with configuration required to start installing fish food.
`
)

type initCmd struct {
	clientOnly bool
	dryRun     bool
}

func newInitCmd() *cobra.Command {
	i := &initCmd{}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "sets up local environment to work with Draft",
		Long:  initDesc,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return i.run()
		},
	}

	f := cmd.Flags()
	f.BoolVar(&i.dryRun, "dry-run", false, "go through all the steps without actually installing anything")

	return cmd
}

// runInit initializes local config and installs Draft to Kubernetes Cluster
func (i *initCmd) run() error {
	home := fish.Home(fish.HomePath)
	userHome := fish.UserHome(fish.UserHomePath)
	dirs := []string{
		filepath.Dir(home.String()),
		home.String(),
		home.Barrel(),
		home.Rigs(),
		userHome.Cache(),
	}

	if !i.dryRun {
		if err := ensureDirectories(dirs); err != nil {
			return err
		}
		return ensureFood()
	}
	return nil
}
