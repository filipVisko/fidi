package cmd

import (
	"os"
	"os/exec"

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
	cmd := exec.Command(gitCommand, "fetch", "--all")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		logger.Fatalf("unable to 'git fetch --all': %s\n", err)
	}
}
