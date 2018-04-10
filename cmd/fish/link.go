package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

func newLinkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link <food>",
		Short: "link fish food",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := getFood(args[0])
			if err != nil {
				return err
			}
			pkg := f.GetPackage(runtime.GOOS, runtime.GOARCH)
			if pkg == nil {
				return fmt.Errorf("food '%s' does not support the current platform (%s/%s)", f.Name, runtime.GOOS, runtime.GOARCH)
			}
			return f.Link(pkg)
		},
	}
	return cmd
}
