package main

import (
	"fmt"

	"github.com/fishworks/gofish/pkg/ohai"
	"github.com/spf13/cobra"
)

func newLintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lint <food>",
		Short: "lint fish food",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := getFood(args[0])
			if err != nil {
				return err
			}
			errs := f.Lint()
			for _, err := range errs {
				ohai.Warningln(err)
			}
			if len(errs) != 0 {
				return fmt.Errorf("%d errors encountered while linting %s", len(errs), f.Name)
			}
			return nil
		},
	}
	return cmd
}
