// Rhythmo — Build better habits with rhythmo
// https://github.com/Sherida101/Rhythmo
//
// See LICENSE for copyright details.

// This file defines the root command for the Rhythmo CLI tool.
// Run this with: go run main.go

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var apiURL string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rhythmo",
	Short: "Track your habits and build better routines",
	Long: `Rhythmo is a CLI tool to help track habits and build better routines.
	You can add, list and manage your habits easily from the command line.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	apiURL = os.Getenv("API_URL")
	rootCmd.PersistentFlags().StringVar(&apiURL, "api", apiURL, "Base URL of the API")
}
