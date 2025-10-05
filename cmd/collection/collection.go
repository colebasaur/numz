package collection

import (
	"github.com/spf13/cobra"
)

var CollectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Interact with your collection of numbers",
	Long:  `Interact with your collection of numbers.`,
}
