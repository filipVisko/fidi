package git

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func runCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error running command '%s %s': %s. stderr: %s", name, strings.Join(arg, " "), err.Error(), stderr.String())
	}

	return nil
}

func getCommonDir() (string, error) {
	gitOutput, err := exec.Command("git", "rev-parse", "--git-common-dir").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git exec error: %s: %s", gitOutput, err)
	}
	commonDir := strings.TrimSpace(string(gitOutput))

	// error out if within the main worktree or a linked worktree of a non-bare repo
	if isNotBareRepo(commonDir) {
		return "", fmt.Errorf("current dir is not a worktree of a bare repo")
	}
	return commonDir, nil
}

func isNotBareRepo(commonDir string) bool {
	return strings.Contains(commonDir, string(os.PathSeparator)+".git") || commonDir == ".git"
}

func branchExists(branchName string, commonDir string) (bool, error) {
	_, err := os.Stat(fmt.Sprintf("%s/refs/remotes/origin/%s", commonDir, branchName))
	if err == nil {
		return true, nil
	} else if !os.IsNotExist(err) {
		return false, fmt.Errorf("error checking for branch existance: %s", err)
	}

	// branch does not exist in the refs/remotes/origin directory, check the packed-refs file
	refsFile, err := os.Open(filepath.Join(commonDir, "packed-refs"))
	if err != nil {
		return false, fmt.Errorf("cannot open the repo's packed-refs file: %s", err)
	}
	defer refsFile.Close()

	scanner := bufio.NewScanner(refsFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, fmt.Sprintf("refs/heads/%s", branchName)) {
			return true, nil
		}
	}
	return false, nil
}
