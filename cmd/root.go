/*
Copyright © 2025 Vicky Chhetri <vickychhetri4@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "acli",
	Short: "acli — a simple daily work log manager",
	Long: `acli is a lightweight CLI tool for recording your daily work logs.
			Each task you add is automatically saved into a date-wise CSV file, making it
			easy to track what you worked on throughout the day.

			Use acli to add tasks, list tasks for a specific date, and maintain a clean
			timeline of your work activities. Ideal for developers, professionals, or anyone
			who wants a fast and minimal personal logging system.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
