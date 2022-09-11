package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func AddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Creates a new branch as a worktree inside of a bare repo.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			name := args[0]
			AddWorktree(name, logger)
		},
	}
	return cmd
}

func AddWorktree(name string, logger *logrus.Logger) {
	var cmd *exec.Cmd
	repoPath, err := getRepoPath()
	if err != nil {
		logger.Fatalf("unable to find the bare repo's path")
	}

	path := filepath.Join(repoPath, name)
	cmd = exec.Command(gitCommand, workTree, "add", path)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		logger.Fatalf("unable to add worktree: %s\n", err)
	}

	fmt.Println(path) // evaluate into a bash variable for cd
}
