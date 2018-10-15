package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/semver"
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
			foodNames = findFood(args)

			nothingUpgraded := true
			for _, name := range foodNames {
				// Check the Barrel name for version
				foodInfo := strings.Split(name, "/")
				nameVer := foodInfo[len(foodInfo)-1]
				if len(foodInfo) == 1 {
					nameVer = foodInfo[0]
				}
				installedVersions := findFoodVersions(nameVer)
				vs := make(semver.Collection, len(installedVersions))
				for i, r := range installedVersions {
					v, err := semver.NewVersion(r)
					if err != nil {
						return fmt.Errorf("Error parsing version: %v", err)
					}
					vs[i] = v
				}
				// we can safely assume there's at least one release installed
				latestInstalledVersion := vs[len(vs)-1]
				food, err := getFood(name)
				if err != nil {
					return err
				}
				if latestInstalledVersion.String() == food.Version {
					continue
				}
				nothingUpgraded = false
				ohai.Ohaif("Upgrading %s...\n", food.Name)
				start := time.Now()
				if err := food.Install(); err != nil {
					return err
				}
				t := time.Now()
				ohai.Successf("%s %s: upgraded in %s\n", food.Name, food.Version, t.Sub(start).String())
			}
			if nothingUpgraded {
				ohai.Successf("Everything up to date!\n")
			}
			return nil
		},
	}
	return cmd
}
