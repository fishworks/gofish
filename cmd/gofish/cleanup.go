package main

import (
	"github.com/fishworks/gofish"
	"github.com/spf13/cobra"
)

func newCleanupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cleanup",
		Short: "cleanup unlinked fish food",
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, name := range findFood([]string{}) {
				versions := findFoodVersions(name)
				if len(versions) > 1 {
					for _, ver := range versions {
						f := gofish.Food{
							Name:    name,
							Version: ver,
						}
						if !f.Linked() {
							if err := f.Uninstall(); err != nil {
								return err
							}
						}
					}
				}
			}
			return nil
		},
	}
	return cmd
}
