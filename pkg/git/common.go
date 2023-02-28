package git

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

func getCommonDir() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--git-common-dir").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git exec error: %s: %s", out, err)
	}
	commonDir := string(out)

	// error out if within the main worktree or a linked worktree of a non-bare repo
	if strings.Contains(commonDir, string(os.PathSeparator)+".git") || commonDir == ".git" {
		return "", fmt.Errorf("current dir is not a worktree of a bare repo")
	}
	return strings.TrimSpace(commonDir), nil
}
