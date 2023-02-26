package main

import (
	"os"

	"github.com/filipVisko/fidi/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "fidi",
		Short: "A git wrapper for managing bare repositories.",
	}
	logger := &log.Logger{
		Out: os.Stderr,
		Formatter: &log.TextFormatter{
			DisableTimestamp:       true,
			DisableLevelTruncation: true,
		},
		Hooks:    make(log.LevelHooks),
		Level:    log.InfoLevel,
		ExitFunc: os.Exit,
	}
	rootCmd.AddCommand(
		cmd.AddCommand(logger),
		cmd.CloneCommand(logger),
		cmd.FetchCommand(logger),
		cmd.ForceRemoveCommand(logger),
		cmd.PullCommand(logger),
		cmd.RemoveCommand(logger),
	)

	_ = rootCmd.Execute()
}
