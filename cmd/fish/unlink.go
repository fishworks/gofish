package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

type unlinkCmd struct{}

func newUnlinkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlink <food>",
		Short: "unlink fish food",
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
			return f.Unlink(pkg)
		},
	}
	return cmd
}
