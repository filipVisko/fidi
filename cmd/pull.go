package cmd

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func PullCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pulls remote changes into a worktree.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			name := args[0]
			pullBranch(name, logger)
		},
	}
	return cmd
}

func pullBranch(name string, logger *logrus.Logger) {
	repoPath, err := getRepoPath()
	if err != nil {
		logger.Fatal(err)
	}
	err = os.Chdir(filepath.Join(repoPath, name))
	if err != nil {
		logger.Fatal(err)
	}
	cmd := exec.Command(gitCommand, "pull")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		logger.Fatalf("unable to pull: %s\n", err)
	}
}
