package collection

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show your collection of numbers",
	Long:  `Show your collection of numbers.`,
	Run:   show,
}

func show(cmd *cobra.Command, args []string) {
	fmt.Println("Your collection of numbers is empty.")
}

func init() {
	CollectionCmd.AddCommand(showCmd)
}

