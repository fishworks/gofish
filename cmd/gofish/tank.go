package main

import (
	"fmt"

	"github.com/fishworks/fish"
	"github.com/spf13/cobra"
)

type tank map[string]string

func (t tank) fill() {
	fishHome := fish.Home(fish.HomePath)
	userHome := fish.UserHome(fish.UserHomePath)

	t["FISH_HOME"] = fishHome.String()
	t["FISH_CACHE"] = userHome.Cache()
	t["FISH_BARREL"] = fishHome.Barrel()
	t["FISH_RIGS"] = fishHome.Rigs()
	t["FISH_DEFAULT_RIG"] = fishHome.DefaultRig()
}

func newTankCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tank",
		Short: "display information about fish's environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			t := tank{}
			t.fill()
			for k, v := range t {
				fmt.Printf("export %s=%q\n", k, v)
			}
			fmt.Print("# Run this command to configure your shell:\n# eval $(gofish tank)\n")
			return nil
		},
	}
	return cmd
}
