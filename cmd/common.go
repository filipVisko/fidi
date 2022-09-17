package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	gitCommand = "git"
	bareRepo   = "bare"
	workTree   = "worktree"
)

func getRepoPath() (string, error) {
	var path string
	dirType, err := getDirType()
	if err != nil {
		return "", err
	}
	if dirType == bareRepo {
		path, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("unable to get current wd: %s", err)
		}
	} else if dirType == workTree {
		out, err := exec.Command(gitCommand, "rev-parse", "--git-common-dir").Output()
		path = string(out)
		if err != nil {
			return "", fmt.Errorf("git error: %s", err)
		}
	} else {
		return "", fmt.Errorf("%q is an unknown directory type", dirType)
	}
	if path == ".git" {
		return "", fmt.Errorf("not a bare repo")
	}
	return strings.TrimSpace(path), nil
}

func isBare() (bool, error) {
	r, err := exec.Command(gitCommand, "rev-parse", "--is-bare-repository").Output()
	if err != nil {
		return false, fmt.Errorf("git error: %s", err)
	}
	boolVal, err := strconv.ParseBool(strings.TrimSpace(string(r)))
	if err != nil {
		return false, fmt.Errorf("unable to parse return value from git: %s", err)
	}
	return boolVal, nil
}

func isWorkingTree() (bool, error) {
	r, err := exec.Command(gitCommand, "rev-parse", "--is-inside-work-tree").Output()
	if err != nil {
		return false, fmt.Errorf("git error: %s", err)
	}
	boolVal, err := strconv.ParseBool(strings.TrimSpace(string(r)))
	if err != nil {
		return false, fmt.Errorf("unable to parse return value from git: %s", err)
	}
	return boolVal, nil
}

func getDirType() (string, error) {
	b, err := isBare()
	if err != nil {
		return "", err
	}
	if b {
		return bareRepo, nil
	}
	w, err := isWorkingTree()
	if err != nil {
		return "", err
	}
	if w {
		return workTree, nil
	}
	return "", fmt.Errorf("not in a git repository")
}
