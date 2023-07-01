package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "An attempt to create a CLI using Cobra",
		Long:  `A command line Ip tracker that tells some info.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
