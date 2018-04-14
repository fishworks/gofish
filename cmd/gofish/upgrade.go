package main

import (
	"time"

	"github.com/fishworks/gofish/pkg/ohai"
	"github.com/spf13/cobra"
)

func newUpgradeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade [food..]",
		Short: "upgrade all fish food. If arguments are provided, only the specified fish foods are upgraded.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := updateRigs(); err != nil {
				return err
			}
			var foodNames []string
			if len(args) > 0 {
				foodNames = args
			} else {
				foodNames = findFood()
			}
			for _, name := range foodNames {
				food, err := getFood(name)
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
			return nil
		},
	}
	return cmd
}
