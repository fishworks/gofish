package main

import (
	"errors"

	"github.com/spf13/cobra"
)

type upgradeCmd struct{}

func newUpgradeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade [food..]",
		Short: "upgrade all fish food. If arguments are provided, only the specified fish foods are upgraded.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}
	return cmd
}
