package main

import (
	"errors"

	"github.com/spf13/cobra"
)

func newTankCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tank",
		Short: "display information about fish's environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("`fish tank` is not implemented")
		},
	}
	return cmd
}
