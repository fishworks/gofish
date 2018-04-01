package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/fishworks/fish"
	"github.com/fishworks/fish/pkg/ohai"
	"github.com/spf13/cobra"
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

const installDesc = `
Install fish food.
`

func newInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install [food]",
		Short: "install fish food",
		Long:  installDesc,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fishFood := args[0]
			if strings.Contains(fishFood, "./\\") {
				return fmt.Errorf("food name '%s' is invalid. Food names cannot include the following characters: './\\'", fishFood)
			}
			l := lua.NewState()
			defer l.Close()
			if err := l.DoFile(filepath.Join(fish.Home(fish.HomePath).DefaultRig(), "Food", fmt.Sprintf("%s.lua", fishFood))); err != nil {
				return err
			}
			var food fish.Food
			if err := gluamapper.Map(l.GetGlobal(strings.ToLower(reflect.TypeOf(food).Name())).(*lua.LTable), &food); err != nil {
				return err
			}
			ohai.Ohailn("Installing draft from github.com/fishworks/fish-food")
			start := time.Now()
			if err := food.Install(); err != nil {
				return err
			}
			t := time.Now()
			ohai.Successf("%s %s: installed in %s\n", food.Name, food.Version, t.Sub(start).String())
			return nil
		},
	}
	return cmd
}
