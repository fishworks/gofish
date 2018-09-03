package main

import (
	"fmt"
	"strings"

	"github.com/fishworks/gofish"

	"github.com/bacongobbler/browser"
	"github.com/spf13/cobra"
)

func newInfoCmd() *cobra.Command {
	var showHomepage bool
	cmd := &cobra.Command{
		Use:   "info <food...>",
		Short: "display information about a particular flavour of fish food",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var b strings.Builder
			var fudz []*gofish.Food
			for _, arg := range args {
				f, err := getFood(arg)
				if err != nil {
					return err
				}
				fudz = append(fudz, f)
			}
			for _, f := range fudz {
				foodVersions := findFoodVersions(f.Name)
				fmt.Fprintf(&b, "%s: ", f.Name)
				if len(foodVersions) == 0 {
					fmt.Fprintln(&b, "no installed versions")
				} else {
					fmt.Fprintf(&b, "%s\n", strings.Join(foodVersions, "\t"))
				}
				if showHomepage {
					if err := browser.Open(f.Homepage); err != nil {
						return err
					}
				}
			}
			fmt.Print(b.String())
			return nil
		},
	}

	p := cmd.Flags()
	p.BoolVar(&showHomepage, "open-homepage", false, "open a browser to the fish food's homepage")

	return cmd
}
