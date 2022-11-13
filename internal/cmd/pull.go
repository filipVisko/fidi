package cmd

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func PullCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pulls remote changes into a worktree.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			pullBranch(name, logger)
		},
	}
	return cmd
}

func pullBranch(name string, logger *log.Logger) {
	commonDir, err := GetCommonDir()
	if err != nil {
		logger.Fatal(err)
	}
	// ideally we should check for existance of worktree first
	err = os.Chdir(filepath.Join(commonDir, name))
	if err != nil {
		logger.Fatal(err)
	}
	err = runCmd("git", "pull")
	if err != nil {
		logger.Fatal(err)
	}
}
