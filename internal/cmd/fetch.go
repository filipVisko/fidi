package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func FetchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Runs 'git fetch --all'",
		Run: func(cmd *cobra.Command, _ []string) {
			logger := logrus.New()
			fetch(logger)
		},
	}
	return cmd
}

func fetch(logger *logrus.Logger) {
	_ = runCmd(gitCommand, "fetch", "--all")
}
