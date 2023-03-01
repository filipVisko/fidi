package git

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"go.uber.org/multierr"
)

// add a new worktree and track any similarly named remote branches
func AddWorktree(name string) error {
	commonDir, err := getCommonDir()
	if err != nil {
		return err
	}

	exists, ok := branchExists(name, commonDir)
	if ok != nil {
		return ok
	}

	args := []string{"worktree", "add", "-b", name, filepath.Join(commonDir, name)}
	if exists {
		args = []string{"worktree", "add", name, filepath.Join(commonDir, name), "--guess-remote"}
	}

	err = runCmd("git", args...)
	if err != nil {
		return err
	}
	fmt.Println(filepath.Join(commonDir, name)) // show the path to allow for easy cd
	return nil
}

// clone repo as bare and configure it to track remote branches
func CloneRepo(url string) error {
	if !strings.Contains(url, ".git") {
		return fmt.Errorf("url must contain a .git suffix")
	}
	err := runCmd("git", "clone", "--bare", url)
	if err != nil {
		return fmt.Errorf("error cloning git repo: %s", err)
	}

	repoPath := path.Base(url)
	err = os.Chdir(repoPath)
	if err != nil {
		return fmt.Errorf("error changing directory into newly cloned git repo: %s", err)
	}

	// configure the bare repo to track all remote branches
	err = runCmd("git", "config", "remote.origin.fetch", "+refs/heads/*:refs/remotes/origin/*")
	if err != nil {
		return fmt.Errorf("error configuring git repo: %s", err)
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
	err = os.Chdir(filepath.Join(commonDir, name))
	if err != nil {
		return fmt.Errorf("error accessing worktree directory: %s", err)
	}
	err = runCmd("git", "pull")
	if err != nil {
		return fmt.Errorf("git pull error: %s", err)
	}
	return nil
}

func RemoveWorktree(name string, force bool) error {
	name = strings.TrimSuffix(name, string(os.PathSeparator))
	args := []string{name}
	if force {
		args = append(args, "--force")
	}

	worktreeArgs := []string{"worktree", "remove"}
	worktreeArgs = append(worktreeArgs, args...)
	err := runCmd("git", worktreeArgs...)

	branchArgs := []string{"branch", "--delete"}
	branchArgs = append(branchArgs, args...)
	err = multierr.Append(err, runCmd("git", branchArgs...))

	if err != nil {
		return err
	}
	return nil
}
