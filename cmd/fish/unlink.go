package main

import (
	"errors"

	"github.com/spf13/cobra"
)

type unlinkCmd struct{}

func newUnlinkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlink",
		Short: "unlink fish food",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}
	return cmd
}
