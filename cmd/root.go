package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/colebasaur/numz/cmd/collection"
)

var rootCmd = &cobra.Command{
	Use:   "numz",
	Short: "A CLI tool to collect and manage numbers!",
	Long:  `numz is a command-line application that allows you collect and manage numbers!`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.AddCommand(collection.CollectionCmd)
}
