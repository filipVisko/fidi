package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RemoveCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes a worktree and its branch reference.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			removeWorktree(name, logger, false)
		},
	}
	return cmd
}

// fidi avoids using flags because the intention is to wrap commands in shell aliases
func ForceRemoveCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "force-remove",
		Short: "Forcibly removes a worktree and its branch reference.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			removeWorktree(name, logger, true)
		},
	}
	return cmd
}

func removeWorktree(name string, logger *log.Logger, force bool) {
	args := []string{name}
	if force {
		args = append(args, "--force")
	}
	worktreeArgs := []string{"worktree", "remove"}
	worktreeArgs = append(worktreeArgs, args...)
	err := runCmd("git", worktreeArgs...)
	if err != nil {
		logger.Warn(err)
	}
	branchArgs := []string{"branch", "--delete"}
	branchArgs = append(branchArgs, args...)
	err = runCmd("git", branchArgs...)
	if err != nil {
		logger.Warn(err)
	}
}
