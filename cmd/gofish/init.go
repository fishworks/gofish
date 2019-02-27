package main

import (
	"github.com/fishworks/gofish/pkg/home"
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
		Short: "sets up local environment to work with fish",
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

func (i *initCmd) run() error {
	h := home.Home(home.HomePath)
	userHome := home.UserHome(home.UserHomePath)
	dirs := []string{
		h.String(),
		h.Barrel(),
		h.Rigs(),
		home.BinPath,
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
