package main

import (
	"errors"

	"github.com/spf13/cobra"
)

type rigInstallCmd struct{}

func newRigInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "install rigs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}
	return cmd
}
