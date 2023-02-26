package main

import (
	"os"

	"github.com/filipVisko/fidi/pkg/git"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "fidi",
		Short: "A git wrapper for managing bare repositories.",
	}
	logger := &log.Logger{
		Out: os.Stderr,
		Formatter: &log.TextFormatter{
			DisableTimestamp:       true,
			DisableLevelTruncation: true,
		},
		Hooks:    make(log.LevelHooks),
		Level:    log.InfoLevel,
		ExitFunc: os.Exit,
	}
	rootCmd.AddCommand(
		AddCommand(logger),
		CloneCommand(logger),
		FetchCommand(logger),
		ForceRemoveCommand(logger),
		PullCommand(logger),
		RemoveCommand(logger),
	)

	_ = rootCmd.Execute()
}

func AddCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new worktree.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			err := git.AddWorktree(name)
			if err != nil {
				logger.Fatal(err)
			}
		},
	}
	return cmd
}

func CloneCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone a remote repo as bare.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			err := git.CloneRepo(url)
			if err != nil {
				logger.Fatal(err)
			}
		},
	}
	return cmd
}

func PullCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pulls remote changes into a worktree.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			err := git.PullBranch(name)
			if err != nil {
				logger.Fatal(err)
			}
		},
	}
	return cmd
}

func FetchCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Runs 'git fetch --all'",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, _ []string) {
			err := git.Fetch()
			if err != nil {
				logger.Fatal(err)
			}
		},
	}
	return cmd
}

func RemoveCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes a worktree and its branch reference.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			err := git.RemoveWorktree(name, true)
			if err != nil {
				logger.Warn(err)
			}
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
			err := git.RemoveWorktree(name, true)
			if err != nil {
				logger.Warn(err)
			}
		},
	}
	return cmd
}
