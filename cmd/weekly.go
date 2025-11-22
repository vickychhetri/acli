/*
Copyright © 2025 Vicky Chhetri <vickychhetri4@gmail.com>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

// weeklyCmd represents the weekly command
var weeklyCmd = &cobra.Command{
	Use:   "weekly",
	Short: "Show summary of last 7 days",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Weekly Summary (Last 7 Days)")
		fmt.Println("--------------------------------")
		for i := 0; i < 7; i++ {
			day := time.Now().AddDate(0, 0, -i)
			filename := filepath.Join("logs", day.Format("2006-01-02")+".csv")

			file, err := os.Open(filename)
			if err != nil {
				continue
			}
			defer file.Close()

			reader := csv.NewReader(file)
			rows, _ := reader.ReadAll()
			fmt.Printf("%s - %d tasks \n", day.Format("2006-01-02"), len(rows))

		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(weeklyCmd)
}
