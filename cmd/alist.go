/*
Copyright © 2025 Vicky Chhetri <vickychhetri4@gmail.com>
*/
package cmd

import (
	"acli/util"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// alistCmd represents the alist command
var alistCmd = &cobra.Command{
	Use:   "alist",
	Short: "Master list of all recorded logs",
	Long:  `Displays the master list of all recorded logs in a colored and readable format.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		dataDir, _ := util.GetDataDir()
		filename := filepath.Join(dataDir, "Master_Log_Sheet.csv")

		file, err := os.Open(filename)
		if err != nil {
			color.Red("No master logs found")
			return nil
		}
		defer file.Close()

		reader := csv.NewReader(file)
		rows, err := reader.ReadAll()
		if err != nil {
			color.Red("Failed to read master log sheet")
			return err
		}

		if len(rows) == 0 {
			color.Yellow("Master log is empty")
			return nil
		}

		// Header
		title := color.New(color.FgGreen, color.Bold)
		title.Println("\n📌 Master Log Records")
		fmt.Println(color.HiBlackString("-----------------------------------------------------------"))

		for _, row := range rows {
			dateStr := color.CyanString(row[0])
			timeStr := color.CyanString(row[1])
			task := row[2]

			// Priority color coding
			var priorityColor string
			switch row[3] {
			case "high":
				priorityColor = color.New(color.FgRed, color.Bold).Sprintf("high")
			case "medium":
				priorityColor = color.New(color.FgYellow).Sprintf("medium")
			case "low":
				priorityColor = color.New(color.FgGreen).Sprintf("low")
			default:
				priorityColor = row[3]
			}

			category := color.MagentaString(row[4])

			// Output format
			fmt.Printf("[%s %s] %s\n", dateStr, timeStr, task)
			fmt.Printf("    Priority: %s   |   Category: %s\n\n", priorityColor, category)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(alistCmd)
}
