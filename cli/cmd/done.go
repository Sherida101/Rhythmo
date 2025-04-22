// Rhythmo — Build better habits with rhythm
// https://github.com/Sherida101/Rhythmo
//
// See LICENSE for copyright details.

package cmd

import (
	"fmt"
	"time"

	"github.com/Sherida101/Rhythmo/cli/db"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done [id]",
	Short:   "Mark a habit as completed",
	Example: `rhythmo done 1 --date 01-10-2025`,
	Long: `Allow the user to mark a habit as done for a specific date.
		This command updates the status of a habit to "done" for the specified date.
		You can also specify a date using the --date flag. If no date is provided,
		the current date will be used.`,
	Args: cobra.ExactArgs(1), // Enforces exactly one positional argument
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		date, _ := cmd.Flags().GetString("date")

		// If no date is provided, use the current date
		if date == "" {
			date = time.Now().Format("02-01-2025") // Format as DD-MM-YYYY
		}

		fmt.Printf("Habit ID %s marked as completed for date %s\n", id, date)

		// Update the habit
		err := db.MarkHabitAsCompleted(id, date)
		if err != nil {
			fmt.Println("❌ Error marking habit as completed:", err)
			return
		}
		fmt.Println("✅ Habit marked as completed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	doneCmd.Flags().StringP("id", "i", "", "ID of the habit to mark as completed")
	doneCmd.MarkFlagRequired("id")

	doneCmd.Flags().StringP("date", "d", "", "Date to mark the habit as completed (DD-MM-YYYY)")
}
