package main

import (
	"fmt"
	"os"
	"task_tracker/app/cmd"
	"task_tracker/app/cmd/mark"
	cmdname "task_tracker/cmd"
)

func init() {
	cmdname.Interlize()
}

func main() {

	switch os.Args[1] {
	case "Add":
		cmd.HandleAdd(os.Args[2:])
	case "Update":
		cmd.HandleUpdate(os.Args[2:])
	case "Delete":
		cmd.HandleDelete(os.Args[2:])
	case "Mark":
		mark.HandleMark(os.Args[2:])
	default:
		fmt.Println("please use right flag")
	}
}
