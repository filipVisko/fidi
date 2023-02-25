package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func GetCommonDir() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--git-common-dir").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git exec error: %s: %s", out, err)
	}
	commonDir := string(out)

	// check the edge cases of being within the main or a linked worktree of a standard repo
	// since fidi cannot handle $GIT_DIR being a child directory in a worktree
	if strings.Contains(commonDir, "/.git") || commonDir == ".git" {
		return "", fmt.Errorf("current dir is not a worktree of a bare repo")
	}
	return strings.TrimSpace(commonDir), nil
}
