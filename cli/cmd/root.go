package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is root command
var RootCmd = &cobra.Command{
	Use: "jmerge-cli",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(MergeCmd)
}
