package main

import (
	"errors"

	"github.com/spf13/cobra"
)

type rigUninstallCmd struct{}

func newRigUninstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall rigs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}
	return cmd
}
