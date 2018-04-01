package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/fishworks/fish"
)

const rigDesc = `
List all installed rigs.
`

func newRigCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rig",
		Short: "list installed rigs",
		Long:  rigDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(fish.Home(fish.HomePath).Rigs())
			return nil
		},
	}
	return cmd
}
