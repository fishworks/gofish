package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fishworks/gofish"
	"github.com/gosuri/uitable"
	"github.com/renstrom/fuzzysearch/fuzzy"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "perform a fuzzy search of fish food",
		Run: func(cmd *cobra.Command, args []string) {
			foundFood := search(args)
			table := uitable.New()
			table.AddRow("NAME", "RIG", "VERSION")
			for _, food := range foundFood {
				f, rig, err := getFood(food)
				if err == nil {
					table.AddRow(f.Name, rig, f.Version)
				} else {
					log.Debugln(err)
				}
			}
			fmt.Println(table)
		},
	}
	return cmd
}

func search(keywords []string) []string {
	var foodNames = findFishFood()
	var foundFood = make(map[string]bool)
	// if no keywords are given, display all available fish food
	if len(keywords) == 0 {
		for _, found := range foodNames {
			foundFood[found] = true
		}
	} else {
		for _, keyword := range keywords {
			for _, found := range fuzzy.Find(keyword, foodNames) {
				foundFood[found] = true
			}
		}
	}
	names := []string{}
	for n := range foundFood {
		names = append(names, n)
	}
	return names
}

func findFishFood() []string {
	home := gofish.Home(gofish.HomePath)
	rigPath := home.Rigs()
	var fudz []string
	filepath.Walk(rigPath, func(p string, f os.FileInfo, err error) error {
		if err != nil {
			log.Errorln(err)
			return nil
		}
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".lua") {
			foodName := strings.TrimSuffix(f.Name(), ".lua")
			repoName := strings.TrimPrefix(p, rigPath+"/")
			repoName = strings.TrimSuffix(repoName, "/"+path.Join("Food", f.Name()))
			name := foodName
			if repoName != home.DefaultRig() {
				name = path.Join(repoName, foodName)
			}
			fudz = append(fudz, name)
		}
		return nil
	})
	sort.Strings(fudz)
	return fudz
}
