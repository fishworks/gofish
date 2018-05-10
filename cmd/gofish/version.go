package main

import (
	"fmt"

	"github.com/fishworks/gofish/version"
	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "display version information",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(version.String())
		},
	}
	return cmd
}
