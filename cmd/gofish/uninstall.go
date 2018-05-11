package main

import (
	"fmt"
	"strings"
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
			relevantFood := search([]string{fishFood})
			switch len(relevantFood) {
			case 0:
				return fmt.Errorf("no fish food with the name '%s' was found", fishFood)
			case 1:
				fishFood = relevantFood[0]
			default:
				// check if we have an exact match
				for _, f := range relevantFood {
					if strings.Compare(f, fishFood) == 0 {
						fishFood = f
						break
					}
				}
				return fmt.Errorf("%d fish food with the name '%s' was found: %v", len(relevantFood), fishFood, relevantFood)
			}
			food, _, err := getFood(fishFood)
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
