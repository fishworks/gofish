package main

import (
	"errors"

	"github.com/spf13/cobra"
)

func newRottenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rotten",
		Short: "show fish food past their best before date (outdated)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("`gofish rotten` is not implemented")
		},
	}
	return cmd
}
