package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "kubectl-command-cli",
		Short: "A cli to find commands related to kubectl",
		Long:  `A cli to quickly find kubectl commands and documentation`,
	}
)

func init() {
	rootCmd.AddCommand(kubesearch)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
