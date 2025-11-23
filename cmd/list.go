/*
Copyright © 2025
*/
package cmd

import (
	"acli/util"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listDate string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List daily logs",
	Long:  `List daily logs with colored output and priority badges.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var logDate time.Time
		var err error

		// ---- Parse date ----
		if listDate == "" {
			logDate = time.Now()
		} else {
			logDate, err = time.Parse("2006-01-02", listDate)
			if err != nil {
				color.Red(" Invalid date. Use YYYY-MM-DD")
				return err
			}
		}

		// ---- File path ----
		dataDir, _ := util.GetDataDir()
		filename := filepath.Join(dataDir, logDate.Format("2006-01-02")+".csv")

		file, err := os.Open(filename)
		if err != nil {
			color.Red(" No logs found for %s", logDate.Format("2006-01-02"))
			return nil
		}
		defer file.Close()

		reader := csv.NewReader(file)
		rows, err := reader.ReadAll()
		if err != nil {
			color.Red(" Failed to read log file")
			return err
		}

		if len(rows) == 0 {
			color.Yellow("No entries recorded for %s", logDate.Format("2006-01-02"))
			return nil
		}

		// ---- Colored headers ----
		title := color.New(color.FgGreen, color.Bold)
		title.Printf("\n📅 Logs for %s\n", logDate.Format("2006-01-02"))
		fmt.Println(color.HiBlackString("----------------------------------------"))

		// ---- Print rows ----
		for i, row := range rows {

			timeStr := row[0]
			task := row[1]
			priority := row[2]
			category := row[3]

			// Priority badge colors
			var priorityText string
			switch priority {
			case "high":
				priorityText = color.New(color.FgRed, color.Bold).Sprintf("high")
			case "medium":
				priorityText = color.New(color.FgYellow, color.Bold).Sprintf("medium")
			case "low":
				priorityText = color.New(color.FgGreen).Sprintf("low")
			default:
				priorityText = priority
			}

			categoryText := color.MagentaString(category)
			timeText := color.CyanString(timeStr)

			fmt.Printf(
				"%d. [%s] %s\n   %s   |   🏷️ %s\n\n",
				i+1,
				timeText,
				task,
				priorityText,
				categoryText,
			)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVar(&listDate, "date", "", "log date (YYYY-MM-DD)")
}
