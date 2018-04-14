// +build windows

package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newTankCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tank",
		Short: "display information about fish's environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			t := tank{}
			t.fill()
			for k, v := range t {
				fmt.Printf("$env:%s = %q\n", k, v)
			}
			fmt.Print("# Run this command to configure your shell:\n# gofish tank | iex\n")
			return nil
		},
	}
	return cmd
}
