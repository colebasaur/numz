package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"math/rand/v2"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a pack of numbers",
	Long:  `Open a pack of numbers.`,
	Run:   Open,
}

func Open(cmd *cobra.Command, args []string) {
	for range 5 {
		fmt.Println(rand.IntN(10) + 1)
	}
}

func init() {
	rootCmd.AddCommand(openCmd)
}
