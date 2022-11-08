package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func AddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new worktree.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			name := args[0]
			addWorktree(name, logger)
		},
	}
	return cmd
}

func addWorktree(name string, logger *logrus.Logger) {
	repoPath, err := getRepoPath()
	if err != nil {
		logger.Fatalf("unable to find the bare repo's path")
	}

	path := filepath.Join(repoPath, name)
	err = runCmd(gitCommand, workTree, "add", path)
	if err != nil {
		logger.Fatalf("unable to add worktree %q: %s\n", path, err)
	}
	// useful to wrap into a shell function to auto-cd into new worktree
	fmt.Println(path)
}
