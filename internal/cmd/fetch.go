package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func FetchCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Runs 'git fetch --all'",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, _ []string) {
			fetch(logger)
		},
	}
	return cmd
}

func fetch(logger *log.Logger) {
	err := runCmd("git", "fetch", "--all")
	if err != nil {
		logger.Fatal(err)
	}
}
