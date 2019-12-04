package main

import (
	"fmt"
	"path/filepath"

	"github.com/fishworks/gofish/pkg/home"
	"github.com/spf13/cobra"
)

func newRigPathCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "path <rig>",
		Short: "display path to a rig",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Println(filepath.Join(home.Rigs(), name))
		},
	}
	return cmd
}
