package main

import (
	"errors"

	"github.com/spf13/cobra"
)

func newSwitchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "switch <food> <version>",
		Short: "switch fish food to another version",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("`gofish switch` is not implemented")
		},
	}
	return cmd
}
