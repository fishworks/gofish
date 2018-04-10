package main

import (
	"errors"

	"github.com/spf13/cobra"
)

func newCleanupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cleanup",
		Short: "cleanup unlinked fish food",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("`fish cleanup` is not implemented")
		},
	}
	return cmd
}
