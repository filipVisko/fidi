package cmd

import (
	"fmt"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func AddCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new worktree.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			addWorktree(name, logger)
		},
	}
	return cmd
}

func addWorktree(name string, logger *log.Logger) {
	repoPath, err := GetCommonDir()
	if err != nil {
		logger.Fatal(err)
	}
	path := filepath.Join(repoPath, name)
	err = runCmd("git", "worktree", "add", path)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(path) // show where we've added a worktree
}
