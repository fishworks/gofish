package main

import (
	"github.com/spf13/cobra"
)

type unlinkCmd struct{}

func newUnlinkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlink",
		Short: "unlink fish food",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := getFood(args[0])
			if err != nil {
				return err
			}
			return f.Unlink()
		},
	}
	return cmd
}
