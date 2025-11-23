/*
Copyright © 2025 Vicky Chhetri <vickychhetri4@gmail.com>
*/
package cmd

import (
	"acli/util"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	dateFlag     string
	priorityFlag string
	categoryFlag string
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a daily log entry",
	Long:  "Add a daily log entry with date, category, and priority.",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			color.Red("Please provide a task:  acli add \"your task\"")
			return errors.New("task required")
		}
		task := args[0]

		var logDate time.Time
		var err error
		if dateFlag == "" {
			logDate = time.Now()
		} else {
			logDate, err = time.Parse("2006-01-02", dateFlag)
			if err != nil {
				color.Red("Invalid date: Use YYYY-MM-DD")
				return err
			}
		}

		dataDir, _ := util.GetDataDir()
		dayFile := filepath.Join(dataDir, logDate.Format("2006-01-02")+".csv")
		masterFile := filepath.Join(dataDir, "Master_Log_Sheet.csv")

		dayF, err := os.OpenFile(dayFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			color.Red("Unable to open daily log file")
			return err
		}
		defer dayF.Close()

		masterF, err := os.OpenFile(masterFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			color.Red("Unable to open master log file")
			return err
		}
		defer masterF.Close()

		timeStr := time.Now().Format("15:04")

		dayWriter := csv.NewWriter(dayF)
		masterWriter := csv.NewWriter(masterF)

		dayWriter.Write([]string{timeStr, task, priorityFlag, categoryFlag})
		masterWriter.Write([]string{
			logDate.Format("2006-01-02"),
			timeStr,
			task,
			priorityFlag,
			categoryFlag,
		})

		dayWriter.Flush()
		masterWriter.Flush()

		green := color.New(color.FgGreen).Add(color.Bold)
		blue := color.New(color.FgCyan)
		magenta := color.New(color.FgMagenta)
		yellow := color.New(color.FgYellow).Add(color.Bold)

		green.Println("\n✓ Task Added Successfully!")

		fmt.Printf("   [Task] %s\n", task)
		fmt.Printf("   [Date] %s  %s\n",
			blue.Sprintf("%s", logDate.Format("2006-01-02")),
			blue.Sprintf("%s", timeStr),
		)
		fmt.Printf("   Category: %s\n", magenta.Sprintf("%s", categoryFlag))
		fmt.Printf("   Priority: %s\n\n", yellow.Sprintf("%s", priorityFlag))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&dateFlag, "date", "d", "", "log date (YYYY-MM-DD)")
	addCmd.Flags().StringVarP(&priorityFlag, "priority", "p", "medium", "task priority (low/medium/high)")
	addCmd.Flags().StringVarP(&categoryFlag, "category", "c", "general", "task category")
}
