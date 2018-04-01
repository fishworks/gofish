package main

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// logLevel is a value to indicate how verbose the user would like the logs to be.
	logLevel   int
)

var globalUsage = `The package manager.
`

func newRootCmd(out io.Writer, in io.Reader) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "fish",
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
		newHomeCmd(out),
		newInitCmd(out),
		newInstallCmd(out),
		newRigCmd(out),
	)

	return cmd
}

func main() {
	cmd := newRootCmd(os.Stdout, os.Stdin)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
