// Rhythmo — Build better habits with rhythm
// https://github.com/Sherida101/Rhythmo
//
// See LICENSE for copyright details.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a habit",
	Long:    `Delete a habit from the list. You can specify the habit ID to delete.`,
	Example: `rhythmo delete 1`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().StringP("id", "i", "", "ID of the habit to delete")
	deleteCmd.MarkFlagRequired("id")
}
