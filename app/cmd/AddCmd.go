package cmd

import (
	"fmt"
	task "task_tracker/app/models/Task"
	"task_tracker/pkg/count"
	"time"
)

func HandleAdd(args []string) {
	taskName := args[0]
	description := args[1]

	//初始化任务模型
	var taskModel = task.Task{
		Id:         count.Count() + 1,
		TaskName:   taskName,
		Decription: description,
		Status:     "todo",
		CreatedAt:  time.Now().Format("2006-01-02"),
		UpdatedAt:  time.Now().Format("2006-01-02"),
	}

	//将任务写入JSON
	err := taskModel.WriteToJSON()
	if err != nil {
		fmt.Println("failed to write to json: ", err)
	}
}
