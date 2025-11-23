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

// weeklyCmd represents the weekly command
var weeklyCmd = &cobra.Command{
	Use:   "weekly",
	Short: "Show summary of last 7 days",
	Long:  `Displays a 7-day summary with dates and number of tasks recorded each day.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		header := color.New(color.FgGreen, color.Bold)
		header.Println("\nWeekly Summary (Last 7 Days)")
		fmt.Println(color.HiBlackString("------------------------------------------"))

		dataDir, _ := util.GetDataDir()
		anyData := false

		for i := 0; i < 7; i++ {
			day := time.Now().AddDate(0, 0, -i)
			dateStr := day.Format("2006-01-02")
			filepath := filepath.Join(dataDir, dateStr+".csv")

			file, err := os.Open(filepath)
			if err != nil {
				continue
			}

			reader := csv.NewReader(file)
			rows, err := reader.ReadAll()
			file.Close()

			if err != nil {
				continue
			}

			anyData = true

			// Colored output
			dateColor := color.New(color.FgCyan, color.Bold).Sprintf(dateStr)
			countColor := color.New(color.FgYellow, color.Bold).Sprintf("%d", len(rows))

			fmt.Printf("%s  -  %s tasks\n", dateColor, countColor)
		}

		if !anyData {
			color.Red("No logs found for the last 7 days.")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(weeklyCmd)
}
