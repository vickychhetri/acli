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

var (
	date        string
	addPriority string
	addCategory string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a daily log entry",
	Long:  `Add a daily log entry`,
	RunE: func(cmd *cobra.Command, args []string) error {
		task := args[0]
		var logDate time.Time
		var err error
		if date == "" {
			logDate = time.Now()
		} else {
			logDate, err = time.Parse("2006-01-02", date)
			if err != nil {
				return fmt.Errorf("invalid date, use YYYY-MM-DD")
			}
		}
		folder := "logs"
		os.MkdirAll(folder, 0755)
		filename := filepath.Join(folder, logDate.Format("2006-01-02")+".csv")
		filenameMaster := filepath.Join(folder, "Master_Log_Sheet.csv")

		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		fileMaster, err := os.OpenFile(filenameMaster, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer fileMaster.Close()

		//date wise file
		writer := csv.NewWriter(file)
		defer writer.Flush()
		timeStr := time.Now().Format("15:04")
		writer.Write([]string{timeStr, task, addPriority, addCategory})

		//Added in master file
		writerMaster := csv.NewWriter(fileMaster)
		defer writerMaster.Flush()
		writerMaster.Write([]string{logDate.Format("2006-01-02"), timeStr, task, addPriority, addCategory})

		fmt.Println("Task Added:", task)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&date, "date", "", "log date (YYYY-MM-DD)")
	addCmd.Flags().StringVar(&addPriority, "priority", "medium", "task priority (low/medium/high)")
	addCmd.Flags().StringVar(&addCategory, "category", "general", "task category")
}
