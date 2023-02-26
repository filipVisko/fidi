package git

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func AddWorktree(name string) error {
	commonDir, err := getCommonDir()
	if err != nil {
		return err
	}
	args := []string{"worktree", "add", filepath.Join(commonDir, name)}

	// if remote branch exists, track it
	_, err = os.Stat(fmt.Sprintf("%s/refs/remotes/origin/%s", commonDir, name))
	if err == nil {
		args = append(args, "--track", name)
	}

	err = runCmd("git", args...)
	if err != nil {
		return err
	}
	fmt.Println(filepath.Join(commonDir, name)) // show the new worktree path
	return nil
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
		return fmt.Errorf("unable to track remote refs: %s", err)
	}
	return nil
}

func Fetch() error {
	return runCmd("git", "fetch", "--all")
}

func PullBranch(name string) error {
	commonDir, err := getCommonDir()
	if err != nil {
		return err
	}
	// ideally we should check for existance of worktree first
	err = os.Chdir(filepath.Join(commonDir, name))
	if err != nil {
		return err
	}
	err = runCmd("git", "pull")
	if err != nil {
		return err
	}
	return nil
}

func RemoveWorktree(name string, force bool) error {
	args := []string{name}
	if force {
		args = append(args, "--force")
	}
	worktreeArgs := []string{"worktree", "remove"}
	worktreeArgs = append(worktreeArgs, args...)
	err := runCmd("git", worktreeArgs...)
	if err != nil {
		return err
	}
	branchArgs := []string{"branch", "--delete"}
	branchArgs = append(branchArgs, args...)
	err = runCmd("git", branchArgs...)
	if err != nil {
		return err
	}
	return nil
}
