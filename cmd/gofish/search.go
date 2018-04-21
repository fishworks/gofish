package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fishworks/gofish"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
)

func newSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "perform a fuzzy search of fish food",
		Run: func(cmd *cobra.Command, args []string) {
			var foodNames = findFishFood()
			var foundFood = make(map[string]bool)
			if len(args) == 0 {
				for _, found := range foodNames {
					foundFood[found] = true
				}
			} else {
				for _, keyword := range args {
					for _, found := range fuzzy.Find(keyword, foodNames) {
						foundFood[found] = true
					}
				}
			}
			var foodList []string
			for food := range foundFood {
				foodList = append(foodList, food)
			}
			fmt.Println(strings.Join(foodList, "\t\t\t"))
		},
	}
	return cmd
}

func findFishFood() []string {
	rigPath := gofish.Home(gofish.HomePath).Rigs()
	var fudz []string
	filepath.Walk(rigPath, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(f.Name(), ".lua") {
			fudz = append(fudz, strings.TrimSuffix(f.Name(), ".lua"))
		}
		return nil
	})
	return fudz
}
