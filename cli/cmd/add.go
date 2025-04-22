// Rhythmo — Build better habits with rhythm
// https://github.com/Sherida101/Rhythmo
//
// See LICENSE for copyright details.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new habit",
	Long: `Add a new habit to your list.
		This command creates a new habit and set its initial parameters.`,
	Example: `rhythmo add "Exercise"
	--frequency daily
	--start-date 01-10-2025 --end-date 01-10-2025
	--description "Daily workout"
	--category health
	--streak 0
	--status active
	--reminder "daily at 9 AM"
	--priority high`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here define the flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().String("frequency", "f", "Frequency of the habit (e.g., daily, weekly)")
	addCmd.Flags().String("start-date", "s", "Start date for the habit (YYYY-MM-DD)")
	addCmd.Flags().String("end-date", "e", "End date for the habit (YYYY-MM-DD)")
	addCmd.Flags().String("description", "desc", "Description of the habit")
	addCmd.Flags().String("category", "c", "Category of the habit (e.g., health, productivity)")
	addCmd.Flags().String("streak", "st", "Streak of the habit (number of days)")
	addCmd.Flags().String("status", "status", "Status of the habit (e.g., active, inactive)")
	addCmd.Flags().String("reminder", "r", "Reminder for the habit (e.g., daily at 9 AM)")
	addCmd.Flags().String("priority", "p", "Priority of the habit (e.g., high, medium, low)")
}
