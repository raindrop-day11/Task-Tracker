package count

import (
	"encoding/json"
	"os"
	task "task_tracker/app/models/Task"
)

func Count() int64 {
	var Tasks []task.Task
	file, _ := os.OpenFile("task.json", os.O_RDONLY, 0644)

	err := json.NewDecoder(file).Decode(&Tasks)
	if err != nil {
		if err.Error() == "EOF" {
			return 0
		}
		os.Exit(1)
	}
	file.Close()
	return Tasks[len(Tasks)-1].Id
}
