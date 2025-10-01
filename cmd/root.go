package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	configFile string
	verbose    bool
)

var rootCmd = &cobra.Command{
	Use:   "numz",
	Short: "A CLI tool to collect and manage numbers",
	Long:  `number-collector is a command-line application that helps you collect and manage numbers efficiently.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to number-collector! Use --help to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Path to configuration file")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}
