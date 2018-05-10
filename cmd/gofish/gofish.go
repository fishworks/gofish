package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// logLevel is a value to indicate how verbose the user would like the logs to be.
	logLevel int
	rootCmd  *cobra.Command
)

var globalUsage = `The package manager.
`

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "gofish",
		Short:        globalUsage,
		Long:         globalUsage,
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			log.SetLevel(log.Level(logLevel))
		},
	}
	p := cmd.PersistentFlags()
	p.IntVar(&logLevel, "log-level", int(log.PanicLevel), "log level")

	cmd.AddCommand(
		newCleanupCmd(),
		newCreateCmd(),
		newHomeCmd(),
		newInfoCmd(),
		newInitCmd(),
		newInstallCmd(),
		newLinkCmd(),
		newListCmd(),
		newPinCmd(),
		newRigCmd(),
		newRottenCmd(),
		newSearchCmd(),
		newSwitchCmd(),
		newTankCmd(),
		newUninstallCmd(),
		newUnlinkCmd(),
		newUnpinCmd(),
		newUpdateCmd(),
		newUpgradeCmd(),
		newVersionCmd(),
	)

	return cmd
}

func main() {
	rootCmd = newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
