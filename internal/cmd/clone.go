package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CloneCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone a remote repo as bare.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			cloneRepo(url, logger)
		},
	}
	return cmd
}

func cloneRepo(url string, logger *log.Logger) {
	err := runCmd("git", "clone", "--bare", url)
	if err != nil {
		logger.Fatal(err)
	}
}
