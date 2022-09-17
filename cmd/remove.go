package cmd

import (
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes a worktree and its branch reference.",
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
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			name := args[0]
			removeWorktree(name, logger, true)
		},
	}
	return cmd
}

func removeWorktree(name string, logger *logrus.Logger, force bool) {
	cmd := exec.Command(gitCommand, "worktree", "remove", name)
	if force {
		cmd = exec.Command(gitCommand, "worktree", "remove", name, "--force")
	}
	err := cmd.Run()
	if err != nil {
		logger.Warnf("unable to remove worktree: %s\n", err)
	}
	cmd = exec.Command(gitCommand, "branch", "--delete", name)
	if force {
		cmd = exec.Command(gitCommand, "branch", "--delete", name, "--force")
	}
	err = cmd.Run()
	if err != nil {
		logger.Warnf("unable to delete branch %q: %s\n", name, err)
	}
}
