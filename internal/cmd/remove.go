package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes a worktree and its branch reference.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			name := args[0]
			removeWorktree(name, logger, false)
		},
	}
	return cmd
}

func ForceRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "force-remove",
		Short: "Forcibly removes a worktree and its branch reference.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			name := args[0]
			removeWorktree(name, logger, true)
		},
	}
	return cmd
}

func removeWorktree(name string, logger *logrus.Logger, force bool) {
	args := []string{name}
	if force {
		args = append(args, "--force")
	}

	worktreeArgs := []string{"worktree", "remove"}
	worktreeArgs = append(worktreeArgs, args...)
	_ = runCmd(gitCommand, worktreeArgs...)

	branchArgs := []string{"branch", "--delete"}
	branchArgs = append(branchArgs, args...)
	_ = runCmd(gitCommand, branchArgs...)
}
