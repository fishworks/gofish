package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fishworks/gofish"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list installed fish food. If an argument is provided, list all installed versions of that fish food",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			table := uitable.New()
			if len(args) == 0 {
				table.AddRow("NAME")
				for _, food := range findFood([]string{}) {
					table.AddRow(food)
				}
			} else {
				table.AddRow("NAME", "VERSION", "LINKED")
				for _, ver := range findFoodVersions(args[0]) {
					f := gofish.Food{
						Name:    args[0],
						Version: ver,
					}
					table.AddRow(f.Name, f.Version, f.Linked())
				}
			}
			fmt.Println(table)
			return nil
		},
	}
	return cmd
}

// When a request comes in for a single food only load that one
func compare(X []os.FileInfo, Y []string) []os.FileInfo {
	var ret []os.FileInfo

	for _, x := range X {
		for _, y := range Y {
			if x.Name() == y {
				ret = append(ret, x)
			}
		}
	}
	return ret
}

func findFood(food []string) []string {
	barrelPath := gofish.Home(gofish.HomePath).Barrel()
	var fudz []string
	files, err := ioutil.ReadDir(barrelPath)
	if len(food) > 0 {
		files = compare(files, food)
	}

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
				fileName := f.Name()
				rigConf, err := ioutil.ReadFile(filepath.Join(barrelPath, fileName) + "/rig.conf")
				if err == nil {
					location := strings.TrimSpace(string(rigConf))
					fileName = strings.Join([]string{location, fileName}, "/")
				}
				fudz = append(fudz, fileName)
			}
		}
	}
	return fudz
}

func findFoodVersions(name string) []string {
	barrelPath := gofish.Home(gofish.HomePath).Barrel()
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
