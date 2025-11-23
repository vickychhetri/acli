/*
Copyright © 2025 Vicky Chhetri <vickychhetri4@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v1.0.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show acli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("acli", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
