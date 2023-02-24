package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CloneCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone a remote repo as bare.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			err := cloneRepo(url)
			if err != nil {
				logger.Fatal(err)
			}
		},
	}
	return cmd
}

func CloneRepo(url string) error {
	if !strings.Contains(url, ".git") {
		return fmt.Errorf("url must contain a .git suffix")
	}
	repoPath := path.Base(url)
	err := runCmd("git", "clone", "--bare", url)
	if err != nil {
		return err
	}
	err = os.Chdir(repoPath)
	if err != nil {
		return err
	}
	// configure the bare repo to track all remote branches
	err = runCmd("git", "config", "remote.origin.fetch", "+refs/heads/*:refs/remotes/origin/*")
	if err != nil {
		return fmt.errorf("unable to track remote refs: %s", err)
	}
}
