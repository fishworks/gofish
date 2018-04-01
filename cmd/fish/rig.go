package main

import "github.com/spf13/cobra"

const rigDesc = `
Manage rigs.
`

func newRigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rig",
		Short: "add rigs",
		Long:  rigDesc,
	}
	cmd.AddCommand(
		newRigInstallCmd(),
		newRigListCmd(),
		newRigUninstallCmd(),
	)
	return cmd
}
