package mark

import (
	"fmt"
	"os"
)

func HandleMark(args []string) {
	if len(args) != 2 {
		fmt.Println("wrong number of parameters")
		os.Exit(1)
	}
	status := args[0]
	idstr := args[1]
	switch status {
	case "in-progress":
		MarkInProgress(idstr)
	case "done":
		MarkDone(idstr)
	default:
		fmt.Println("wrong input.only in-progress or done")
	}
}
