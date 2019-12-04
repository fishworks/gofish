package main

import (
	"fmt"

	"github.com/fishworks/gofish/pkg/home"
	"github.com/spf13/cobra"
)

const homeDesc = `
Display the location of fish's home directory. This is where barrels, cached downloads and rigs
live.
`

func newHomeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "home",
		Short: "print the location of fish's home directory",
		Long:  homeDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(home.String())
			return nil
		},
	}
	return cmd
}
