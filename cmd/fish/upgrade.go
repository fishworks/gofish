package main

import (
	"errors"
	"time"

	"github.com/fishworks/fish/pkg/ohai"
	"github.com/spf13/cobra"
)

type upgradeCmd struct{}

func newUpgradeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade [food..]",
		Short: "upgrade all fish food. If arguments are provided, only the specified fish foods are upgraded.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := updateRigs(); err != nil {
				return err
			}
			if len(args) > 0 {
				for _, arg := range args {
					food, err := getFood(arg)
					if err != nil {
						return err
					}
					ohai.Ohaif("Upgrading %s...\n", food.Name)
					start := time.Now()
					if err := food.Install(); err != nil {
						return err
					}
					t := time.Now()
					ohai.Successf("%s %s: upgraded in %s\n", food.Name, food.Version, t.Sub(start).String())
				}
			} else {
				return errors.New("`fish upgrade` is not implemented")
			}
			return nil
		},
	}
	return cmd
}
