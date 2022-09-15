Fidi is an opinionated yet simplified workflow for git branching.

It's meant to be used as an alias of commonly used git commands for managing braches but replacing the workflow for one that uses [worktrees](https://git-scm.com/docs/git-worktree).
Fidi will clone repos as [bare](https://git-scm.com/docs/git-clone#Documentation/git-clone.txt---bare) which then enables you to checkout new branches as ephemeral worktrees that are subfolders of the bare repo as `bare_repo_name.git/branch_name`.
Fidi leaves the `.git` suffix when cloning to differentiate between bare and non-bare repos.

Worktrees are great. I don't know why this feature is not as well known.
- They allow you to have multiple checked-out branches at the same time.
- They make switching branches easier by leaving your untracked changes unmodified inside of their worktree.
  - You no longer need to use the git stash workflow or commit your unfinished changes.
- You could create distinct per-branch [watchman](https://facebook.github.io/watchman/) configurations to trigger commands on file change.

However, there are some DX issues with using worktrees so fidi exists to make them smoother to use.

Note that bare repos don't have a tracking branch. Some incremental linters and build tools don't like that very much.

## Installing fidi

You can use `go install`

```bash
go install github.com/filipVisko/fidi@latest
```

...or download one of the releases and add it to your `$PATH`.

### Add fidi aliases

Fidi was designed to be used as an alias which reduces the need to pass flags.
You can use it directly but aliases are faster and don't require you to mess with your muscle memory for working with git.
Here's a list of aliases that you could use as an example. I'm following the pattern used by the [git plugin for oh-my-zsh](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/git).

```bash
# If you're using the oh-my-zsh git plugin, your '.zshrc' may contain...
alias g="git"
alias ga="git add"
alias gaa="git add --all"
alias gp="git push"

# You could replace parts of the oh-my-zsh provided alias values with the following...
alias gcl="fidi clone"        # similar to 'git clone'
alias gb="fidi add"           # similar to 'git branch'
alias gbd="fidi remove"       # similar to 'git branch -d'
alias gbD="fidi force-remove" # similar to 'git branch -D'

# 'fidi add' can be evaluated into a cd command via script to make it behave similar to 'git checkout -b'
```
