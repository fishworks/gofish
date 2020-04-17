package main

import (
	"fmt"
	"path"
	"reflect"
	"strings"

	"github.com/fishworks/gofish"
	"github.com/fishworks/gofish/pkg/ohai"
	"github.com/spf13/cobra"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

func newLintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lint <file...>",
		Short: "lint fish food",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				l := lua.NewState()
				defer l.Close()
				if err := l.DoFile(arg); err != nil {
					return err
				}
				var food gofish.Food
				if err := gluamapper.Map(l.GetGlobal(strings.ToLower(reflect.TypeOf(food).Name())).(*lua.LTable), &food); err != nil {
					return err
				}
				errs := food.Lint()

				base := strings.TrimSuffix(path.Base(arg), ".lua")
				if base != food.Name {
					errs = append(errs, fmt.Errorf("File name '%v' must match name in file '%v'", base, food.Name))
				}

				for _, err := range errs {
					ohai.Warningln(err)
				}
				if len(errs) != 0 {
					return fmt.Errorf("%d errors encountered while linting %s", len(errs), food.Name)
				}

				ohai.Successf("No errors discovered in '%v'\n", food.Name)
			}
			return nil
		},
	}
	return cmd
}
