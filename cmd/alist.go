/*
Copyright © 2025 Vicky Chhetri <vickychhetri4@gmail.com>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// alistCmd represents the alist command
var alistCmd = &cobra.Command{
	Use:   "alist",
	Short: "Master List of all the recorded logs ",
	Long:  `Master List of all the recorded logs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := filepath.Join("logs", "Master_Log_Sheet.csv")
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("no logs")
		}
		reader := csv.NewReader(file)
		rows, _ := reader.ReadAll()

		for _, row := range rows {
			fmt.Printf("[%s] [%s] %s | Priority: %s | Category: %s\n", row[0], row[1], row[2], row[3], row[4])
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(alistCmd)
}
