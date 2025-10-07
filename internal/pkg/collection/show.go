package collection

import (
	"fmt"
	"github.com/colebasaur/numz/internal/model"
	"github.com/spf13/cobra"
)

func Show(cmd *cobra.Command, args []string) {
	collectionItems, err := model.CollectionItemsGet()
	if err != nil {
		fmt.Println("Error retrieving collection items:", err)
		return
	}

	if len(collectionItems) == 0 {
		fmt.Println("Your collection is empty.")
		return
	}

	fmt.Println("Your number collection:")
	for _, item := range collectionItems {
		fmt.Printf("%d (%d)\n", item.Number, item.Quantity)
	}
}
