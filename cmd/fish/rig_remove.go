package main

import (
	"errors"

	"github.com/spf13/cobra"
)

type rigRemoveCmd struct{}

func newRigRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove rigs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}
	return cmd
}
