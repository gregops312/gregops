package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const CliName = "gregops"

var rootCmd = &cobra.Command{
	Use:   CliName,
	Short: fmt.Sprintf("%s is a CLI tool built with Cobra", CliName),
	Long: fmt.Sprintf(`%s is a CLI application built with Cobra.

This application provides various commands and utilities.
You can use it to perform different tasks from the command line.`, CliName),
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior when no subcommand is specified
		cmd.Printf("Welcome to %s CLI!\n", CliName)
		cmd.Printf("Use '%s --help' to see available commands.\n", CliName)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrln(err.Error())
		os.Exit(1)
	}
}

func init() {
	// Here you can define your flags and configuration settings.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.%s.yaml)", CliName))
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
