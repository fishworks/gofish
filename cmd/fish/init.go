package main

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/fishworks/fish"
)

const (
	initDesc = `
This command sets up fish with the directories required to work with fish
`
)

type initCmd struct {
	clientOnly bool
	dryRun     bool
	out        io.Writer
}

func newInitCmd(out io.Writer) *cobra.Command {
	i := &initCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "sets up local environment to work with Draft",
		Long:  initDesc,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return i.run()
		},
	}

	f := cmd.Flags()
	f.BoolVar(&i.dryRun, "dry-run", false, "go through all the steps without actually installing anything. Mostly used along with --debug for debugging purposes.")

	return cmd
}

// runInit initializes local config and installs Draft to Kubernetes Cluster
func (i *initCmd) run() error {
	home := fish.Home(fish.HomePath)
	userHome := fish.UserHome(fish.UserHomePath)
	dirs := []string{
		home.String(),
		home.Barrel(),
		home.Rigs(),
		userHome.Cache(),
	}

	if !i.dryRun {
		return ensureDirectories(dirs)
	}
	return nil
}
