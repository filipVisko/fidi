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
		Short: "Add a new worktree.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			name := args[0]
			addWorktree(name, logger)
		},
	}
	return cmd
}

func addWorktree(name string, logger *logrus.Logger) {
	var cmd *exec.Cmd
	repoPath, err := getRepoPath()
	if err != nil {
		logger.Fatalf("unable to find the bare repo's path")
	}

	path := filepath.Join(repoPath, name)
	cmd = exec.Command(gitCommand, workTree, "add", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		logger.Fatalf("unable to add worktree %q: %s\n", path, err)
	}
	fmt.Println(path)
}
