package main

import (
	"fmt"

	"github.com/fishworks/gofish"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

func newRottenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rotten",
		Short: "show fish food past their best before date (outdated)",
		RunE: func(cmd *cobra.Command, args []string) error {
			table := uitable.New()
			table.AddRow("NAME", "VERSION")
			for _, name := range findFood() {
				versions := findFoodVersions(name)
				if len(versions) > 1 {
					for _, ver := range versions {
						f := gofish.Food{
							Name:    name,
							Version: ver,
						}
						if !f.Linked() {
							table.AddRow(f.Name, f.Version)
						}
					}
				}
			}
			fmt.Println(table)
			return nil
		},
	}
	return cmd
}
