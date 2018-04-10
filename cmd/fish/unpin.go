package main

import (
	"errors"

	"github.com/spf13/cobra"
)

func newUnpinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unpin <food>",
		Short: "remove protection from a fish food, allowing fish to install upgrades",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("`fish unpin` is not implemented")
		},
	}
	return cmd
}
