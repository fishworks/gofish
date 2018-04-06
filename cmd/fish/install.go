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
			if err := updateRigs(); err != nil {
				return err
			}
			food, err := getFood(fishFood)
			if err != nil {
				return err
			}
			if len(findFoodVersions(fishFood)) > 0 {
				ohai.Ohaif("%s is already installed. Please use `fish upgrade %s` to upgrade.\n", fishFood, fishFood)
				return nil
			}
			ohai.Ohaif("Installing %s...\n", fishFood)
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

func getFood(name string) (*fish.Food, error) {
	if strings.Contains(name, "./\\") {
		return nil, fmt.Errorf("food name '%s' is invalid. Food names cannot include the following characters: './\\'", name)
	}
	l := lua.NewState()
	defer l.Close()
	if err := l.DoFile(filepath.Join(fish.Home(fish.HomePath).DefaultRig(), "Food", fmt.Sprintf("%s.lua", name))); err != nil {
		return nil, err
	}
	var food fish.Food
	if err := gluamapper.Map(l.GetGlobal(strings.ToLower(reflect.TypeOf(food).Name())).(*lua.LTable), &food); err != nil {
		return nil, err
	}
	return &food, nil
}
