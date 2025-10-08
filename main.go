package main

import (
	"github.com/colebasaur/numz/cmd"
	"github.com/colebasaur/numz/internal/storage"
)

func main() {
	err := storage.Init()
	if err != nil {
		panic(err)
	}
	defer storage.Close()
	cmd.Execute()
}
