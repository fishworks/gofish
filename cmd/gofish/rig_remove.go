package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/fishworks/gofish"
	"github.com/fishworks/gofish/pkg/ohai"
	"github.com/spf13/cobra"
)

func newRigRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove <rig...>",
		Short: "remove rigs",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()
			rigPath := gofish.Home(gofish.HomePath).Rigs()
			rigs := findRigs(rigPath)
			foundRigs := map[string]bool{}
			for _, arg := range args {
				foundRigs[arg] = false
			}
			for _, rig := range rigs {
				for _, arg := range args {
					if rig == arg {
						foundRigs[rig] = true
						if err := os.RemoveAll(filepath.Join(rigPath, rig)); err != nil {
							return err
						}
					}
				}
			}
			t := time.Now()
			for rig, found := range foundRigs {
				if !found {
					ohai.Warningf("rig '%s' was not found in the rig list\n", rig)
				}
			}
			ohai.Successf("rigs uninstalled in %s\n", t.Sub(start).String())
			return nil
		},
	}
	return cmd
}
