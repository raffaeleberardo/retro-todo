package main

import (
	"fmt"
	"os"
	"retro-todo/helper"
)

const (
	add      = "add"
	list     = "list"
	complete = "complete"
	delete   = "delete"
)

func main() {

	arg := os.Args[1]

	switch arg {
	case add:
		helper.HandleAdd()
	case list:
		helper.HandleList()
	case complete:
		helper.HandleComplete()
	case delete:
		helper.HandleDelete()
	default:
		fmt.Println("Unknown action")
	}

	// helper.ReadCsv()
}
