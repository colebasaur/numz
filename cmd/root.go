package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/colebasaur/numz/internal/pkg/collection"
	"github.com/colebasaur/numz/internal/pkg/packs"
)

// Root Command
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

// Collection Subcommands
var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Show and Manage your number collection",
	Long:  `Show and Manage your number collection`,
}

var collectionManageCmd = &cobra.Command{
	Use:   "manage",
	Short: "Manage your number collection",
	Long:  `Manage your number collection`,
	Run:   collection.Manage,
}

var collectionShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show your number collection",
	Long:  `Show your number collection`,
	Run:   collection.Show,
}

// Packs Subcommands
var packsCmd = &cobra.Command{
	Use:   "packs",
	Short: "View and Open number packs",
	Long:  `View and Open number packs`,
}

var packsViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View pack details",
	Long:  `View pack details`,
	Run:   packs.View,
}

var packsOpenCmd = &cobra.Command{
	Use:   "open [pack id]",
	Short: "Open pack and reveal numbers",
	Long:  `Open pack by its ID and reveal the numbers inside the pack`,
	Args:  cobra.ExactArgs(1),
	Run:   packs.Open,
}

func init() {
	// Hide "completions" command
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	// Add collection commands
	rootCmd.AddCommand(collectionCmd)
	collectionCmd.AddCommand(collectionManageCmd)
	collectionCmd.AddCommand(collectionShowCmd)

	// Add packs commands
	rootCmd.AddCommand(packsCmd)
	packsCmd.AddCommand(packsViewCmd)
	packsCmd.AddCommand(packsOpenCmd)
}
