package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fishworks/fish"
	"github.com/spf13/cobra"
)

type listCmd struct{}

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list installed fish food",
		RunE: func(cmd *cobra.Command, args []string) error {
			barrelPath := fish.Home(fish.HomePath).Barrel()
			fishFood := findFood(barrelPath)
			fmt.Println(strings.Join(fishFood, "\t\t\t"))
			return nil
		},
	}
	return cmd
}

func findFood(dir string) []string {
	var food []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}
	}

	for _, f := range files {
		if f.IsDir() {
			food = append(food, f.Name())
		}
	}
	return food
}
