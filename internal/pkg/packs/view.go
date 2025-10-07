package packs

import (
	"fmt"
	"github.com/colebasaur/numz/internal/model"
	"github.com/spf13/cobra"
)

func View(cmd *cobra.Command, args []string) {
	packs, err := model.PacksGet()
	if err != nil {
		fmt.Println("Error retrieving packs:", err)
		return
	}

	if len(packs) == 0 {
		fmt.Println("No packs available.")
		return
	}

	fmt.Println("Available packs:")
	for _, pack := range packs {
		fmt.Println(pack)
	}
}
