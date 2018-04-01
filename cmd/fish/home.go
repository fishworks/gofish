package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/fishworks/fish"
)

const homeDesc = `
Display the location of fish's home directory. This is where barrels,
cached downloads and rigs live.
`

func newHomeCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "home",
		Short: "print the location of fish's home directory",
		Long:  homeDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(out, fish.Home(fish.HomePath))
			return nil
		},
	}
	return cmd
}
