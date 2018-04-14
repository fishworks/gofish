package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fishworks/gofish"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

func newRigListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list rigs",
		RunE: func(cmd *cobra.Command, args []string) error {
			rigPath := gofish.Home(gofish.HomePath).Rigs()
			rigs := findRigs(rigPath)

			table := uitable.New()
			table.AddRow("NAME")
			for _, rig := range rigs {
				table.AddRow(rig)
			}
			fmt.Println(table)
			return nil
		},
	}
	return cmd
}

func findRigs(dir string) []string {
	var rigs []string
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() && f.Name() == "Food" {
			rigName := strings.TrimPrefix(filepath.Dir(path), dir+string(os.PathSeparator))
			rigs = append(rigs, rigName)
		}
		return nil
	})
	return rigs
}
