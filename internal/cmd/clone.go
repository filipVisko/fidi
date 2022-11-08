package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CloneCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone a remote repo as bare.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			url := args[0]
			cloneRepo(url, logger)
		},
	}
	return cmd
}

func cloneRepo(url string, logger *logrus.Logger) {
	_ = runCmd(gitCommand, "clone", "--bare", url)
}
