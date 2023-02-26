Fidi is an opinionated yet simplified workflow for git branching.

It wraps your native `git` client and is meant to be used as a replacement command for managing branches.
It does this by creating [worktrees](https://git-scm.com/docs/git-worktree) as sub-folders of a [bare](https://git-scm.com/docs/git-clone#Documentation/git-clone.txt---bare) repo.
This allows you to have multiple checked-out branches on disk and the ability to swap between them using `cd` without the need to use `git stash` for cleaning up unfinished changes.

Fidi makes it simpler to use worktrees when compared to using the git's worktree commands directly:
- It takes over the management of the checkout path when creating new worktrees by always placing them as a sub-folder of the parent bare repo, providing you with a consistent use pattern and uncluttered file tree.
- It can delete a worktree and its branch reference with a single command.
- Makes it easy to pull branch changes while in the worktree of a different branch. Useful for quickly merging `origin/main` into your dev branch.

```bash
$ pwd
~/repo_name.git/main # path to main branch

# create a new branch
$ fidi add test_branch
Preparing worktree (new branch 'test_branch')
~/repo_name.git/test_branch # path to new branch

# use cd to swap between branches
cd ~/repo_name.git/test_branch
```

Note that you can't have a tracking branch when using bare repos. Incremental linters may not like that.

## Installing fidi

```bash
go install github.com/filipVisko/fidi@latest
```

Or download a release and add it to your `$PATH`.

### Add aliases into your shell

Fidi was designed to be used as an alias which reduces the need to pass flags.
As an example here, I'm following the alias pattern used by the [git plugin](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/git) for oh-my-zsh and replaced the following git aliases with fidi.

```bash
# my git aliases
alias g="git"
alias ga="git add"
alias gaa="git add --all"
alias gp="git push"

# fidi aliases
alias gcl="fidi clone"        # similar to 'git clone' but clones the repo as bare
alias gb="fidi add"           # similar to 'git branch', will create a new worktree as a subfolder of the bare repo as 'repo_name.git/branch_name'
alias gpb="fidi pull"         # will pull changes from the remote into the desired branch without changing directory
alias gbd="fidi remove"       # similar to 'git branch -d', will delete the worktree and the branch reference
alias gbD="fidi force-remove" # similar to 'git branch -D', will force delete the worktree and branch reference
```
