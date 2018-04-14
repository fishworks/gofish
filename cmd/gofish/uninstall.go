package main

import (
	"time"

	"github.com/fishworks/gofish/pkg/ohai"
	"github.com/spf13/cobra"
)

func newUninstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall <food>",
		Short: "uninstall fish food",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fishFood := args[0]
			food, err := getFood(fishFood)
			if err != nil {
				return err
			}
			ohai.Ohaif("Uninstalling %s...\n", fishFood)
			start := time.Now()
			if err := food.Uninstall(); err != nil {
				return err
			}
			t := time.Now()
			ohai.Successf("%s %s: uninstalled in %s\n", food.Name, food.Version, t.Sub(start).String())
			return nil
		},
	}
	return cmd
}
