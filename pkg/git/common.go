package git

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func branchExists(name string, commonDir string) (bool, error) {
	_, err := os.Stat(fmt.Sprintf("%s/refs/remotes/origin/%s", commonDir, name))
	if err == nil {
		return true, nil
	}

	refsFile, err := os.Open(filepath.Join(commonDir, "packed-refs"))
	if err != nil {
		return false, fmt.Errorf("cannot open the repo's packed-refs file: %s", err)
	}
	defer refsFile.Close()

	scanner := bufio.NewScanner(refsFile)
	for scanner.Scan() {
		l := scanner.Text()
		if strings.Contains(l, fmt.Sprintf("refs/heads/%s", name)) {
			return true, nil
		}
	}
	return false, nil
}
