package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/fishworks/fish"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list installed fish food. If an argument is provided, list all installed versions of that fish food",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var output []string
			if len(args) == 0 {
				output = findFood()
			} else {
				output = findFoodVersions(args[0])
			}
			fmt.Println(strings.Join(output, "\t\t\t"))
			return nil
		},
	}
	return cmd
}

func findFood() []string {
	barrelPath := fish.Home(fish.HomePath).Barrel()
	var fudz []string
	files, err := ioutil.ReadDir(barrelPath)
	if err != nil {
		return []string{}
	}

	for _, f := range files {
		if f.IsDir() {
			files, err := ioutil.ReadDir(filepath.Join(barrelPath, f.Name()))
			if err != nil {
				continue
			}
			if len(files) > 0 {
				fudz = append(fudz, f.Name())
			}
		}
	}
	return fudz
}

func findFoodVersions(name string) []string {
	barrelPath := fish.Home(fish.HomePath).Barrel()
	var versions []string
	files, err := ioutil.ReadDir(filepath.Join(barrelPath, name))
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
