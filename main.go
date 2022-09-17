package main

import (
	"github.com/filipVisko/fidi/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "fidi",
		Short: "A git wrapper for managing bare repositories.",
	}

	rootCmd.AddCommand(
		cmd.AddCommand(),
		cmd.CloneCommand(),
		cmd.FetchCommand(),
		cmd.ForceRemoveCommand(),
		cmd.PullCommand(),
		cmd.RemoveCommand(),
	)

	_ = rootCmd.Execute()
}
