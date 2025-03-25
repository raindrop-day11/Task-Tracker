package main

import (
	"fmt"
	"os"
	"task_tracker/app/cmd"
)

func main() {

	switch os.Args[1] {
	case "Add":
		cmd.HandleAdd(os.Args[2:])
	case "Update":
		cmd.HandleUpdate(os.Args[2:])
	case "Delete":
		cmd.HandleDelete(os.Args[2:])
	case "Mark":
		cmd.HandleMark(os.Args[2:])
	case "List":
		cmd.HandleList(os.Args[2:])
	default:
		fmt.Println("please use right flag")
	}
}
