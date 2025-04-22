// Rhythmo — Build better habits with rhythm
// https://github.com/Sherida101/Rhythmo
//
// See LICENSE for copyright details.

package cmd

import (
	"fmt"

	"github.com/Sherida101/Rhythmo/cli/api"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all habits",
	Run: func(cmd *cobra.Command, args []string) {
		// Pass an empty string to use API_URL from .env
		habits, err := api.FetchHabits("")
		// habits, err := api.FetchHabits("http://localhost:8080")

		if err != nil {
			fmt.Println("❌ Error fetching habits:", err)
			return
		}

		if len(habits) == 0 {
			fmt.Println("📭 No habits found.")
			return
		}

		for _, h := range habits {
			fmt.Printf("✅ [%d] %s (Streak: %d days)\n", h.ID, h.Name, h.Streak)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
