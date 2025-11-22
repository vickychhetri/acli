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

var listDate string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List daily logs",
	Long:  `List daily logs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var logDate time.Time
		var err error

		if listDate == "" {
			logDate = time.Now()
		} else {
			logDate, err = time.Parse("2006-01-02", listDate)
			if err != nil {
				return fmt.Errorf("invalid date, use YYYY-MM-DD")
			}
		}

		filename := filepath.Join("logs", logDate.Format("2006-01-02")+".csv")
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("no logs for date: %s", logDate.Format("2006-01-02"))
		}
		defer file.Close()

		reader := csv.NewReader(file)
		rows, _ := reader.ReadAll()

		fmt.Println("Logs for ", logDate.Format("2006-01-02"))
		fmt.Println("----------------------------------------")
		for i, row := range rows {
			fmt.Printf("%d. [%s] %s  | Priority: %s | Category: %s\n",
				i+1, row[0], row[1], row[2], row[3])
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVar(&listDate, "date", "", "log date (YYYY-MM-DD)")
}
