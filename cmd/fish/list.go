package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/fishworks/fish"
	"github.com/spf13/cobra"
)

type listCmd struct{}

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list installed fish food. If an argument is provided, list all installed versions of that fish food",
		RunE: func(cmd *cobra.Command, args []string) error {
			var output []string
			barrelPath := fish.Home(fish.HomePath).Barrel()
			if len(args) == 0 {
				output = findFood(barrelPath)
			} else {
				output = findFoodVersions(barrelPath, args[0])
			}
			fmt.Println(strings.Join(output, "\t\t\t"))
			return nil
		},
	}
	return cmd
}

func findFood(dir string) []string {
	var fudz []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}
	}

	for _, f := range files {
		if f.IsDir() {
			fud := fish.Food{Name: f.Name()}
			if fud.Installed() {
				fudz = append(fudz, f.Name())
			}
		}
	}
	return fudz
}

func findFoodVersions(dir, name string) []string {
	var versions []string
	files, err := ioutil.ReadDir(filepath.Join(dir, name))
	if err != nil {
		return []string{}
	}

	for _, f := range files {
		if f.IsDir() {
			versions = append(versions, f.Name())
		}
	}
	return versions
}
