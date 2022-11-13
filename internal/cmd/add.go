package cmd

import (
	"fmt"
	"os"
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
	commonDir, err := GetCommonDir()
	if err != nil {
		logger.Fatal(err)
	}
	args := []string{"worktree", "add", filepath.Join(commonDir, name)}

	// if remote branch exists, track it
	_, err = os.Stat(fmt.Sprintf("%s/refs/remotes/origin/%s", commonDir, name))
	if err == nil {
		args = append(args, "--track", name)
	}

	err = runCmd("git", args...)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(filepath.Join(commonDir, name)) // show the new worktree path
}
