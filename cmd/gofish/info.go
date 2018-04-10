package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func newInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info <food...>",
		Short: "display information about a particular flavour of fish food",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var b strings.Builder
			for _, arg := range args {
				versions := findFoodVersions(arg)
				fmt.Fprintf(&b, "%s: ", arg)
				if len(versions) == 0 {
					fmt.Fprintln(&b, "no installed versions")
				} else {
					fmt.Fprintf(&b, "%s\n", strings.Join(versions, "\t"))
				}
			}
			fmt.Print(b.String())
			return nil
		},
	}
	return cmd
}
