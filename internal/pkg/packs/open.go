package packs

import (
	"fmt"
	"github.com/colebasaur/numz/internal/model"
	"github.com/spf13/cobra"
	"strconv"
)

func Open(cmd *cobra.Command, args []string) {
	packID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid pack ID:", err)
		return
	}

	pack := model.Pack(packID)
	numbers, err := model.PackOpen(pack)
	if err != nil {
		fmt.Println("Error opening pack:", err)
		return
	}

	fmt.Printf("You opened pack %d and got the following numbers:\n", packID)
	for _, number := range numbers {
		fmt.Println(number)
	}
}
