package main

import (
	"time"

	"github.com/fishworks/gofish"
	"github.com/fishworks/gofish/pkg/ohai"

	"github.com/fishworks/gofish/pkg/rig/installer"
	"github.com/spf13/cobra"
)

func newRigAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <rig>",
		Short: "add rigs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := installer.New(args[0], "", gofish.Home(gofish.HomePath))
			if err != nil {
				return err
			}

			start := time.Now()
			if err := installer.Install(i); err != nil {
				return err
			}
			t := time.Now()
			ohai.Successf("rig constructed in %s\n", t.Sub(start).String())
			return nil
		},
	}
	return cmd
}
