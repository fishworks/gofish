package main

import (
	"errors"

	"github.com/spf13/cobra"
)

type linkCmd struct{}

func newLinkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link",
		Short: "link fish food",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}
	return cmd
}
