package main

import (
	"errors"

	"github.com/spf13/cobra"
)

type infoCmd struct{}

func newInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "display information about a particular flavour of fish food (versions, caveats, etc)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}
	return cmd
}
