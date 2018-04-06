package main

import (
	"errors"

	"github.com/spf13/cobra"
)

type pinCmd struct{}

func newPinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pin <food>",
		Short: "protect a fish food, preventing fish from installing upgrades",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("`fish pin` is not implemented")
		},
	}
	return cmd
}
