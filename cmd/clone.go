package cmd

import (
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CloneCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone a remote repo as bare.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			url := args[0]
			cloneRepo(url, logger)
		},
	}
	return cmd
}

func cloneRepo(url string, logger *logrus.Logger) {
	cmd := exec.Command(gitCommand, "clone", "--bare", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		logger.Fatalf("could not clone: %s\n", err)
	}
}