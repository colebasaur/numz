package collection

import (
	"fmt"

	"github.com/spf13/cobra"
)

var manageCmd = &cobra.Command{
	Use:   "manage",
	Short: "Manage your collection of numbers",
	Long:  `Manage your collection of numbers.`,
	Run:   manage,
}

func manage(cmd *cobra.Command, args []string) {
	fmt.Println("Nothing to manage yet.")
}

func init() {
	CollectionCmd.AddCommand(manageCmd)
}
