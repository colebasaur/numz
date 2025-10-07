package main

import (
	"github.com/colebasaur/numz/cmd"
	"github.com/colebasaur/numz/internal/storage"
)

func main() {
	storage.Init()
	defer storage.Close()
	cmd.Execute()
}
