package cmd

import (
	"encoding/json"
	"os"
	task "task_tracker/app/models/Task"
	"task_tracker/pkg/changejson"
	"task_tracker/pkg/logger"
	"time"
)

func HandleAdd(args []string) {
	if len(args) != 2 {
		logger.WarningExit("wrong number of parameters")
	}
	taskName := args[0]
	description := args[1]

	var tasks []task.Task
	//打开文件
	file, err := os.OpenFile("task.json", os.O_CREATE|os.O_RDWR, 0644)
	logger.WarningExitIF("json file can not open", err)
	defer file.Close()

	//解码
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil && err.Error() != "EOF" {
		logger.ErrorExit(err)
	}

	//初始化任务模型
	var taskModel = task.Task{
		Id:         tasks[len(tasks)-1].Id + 1,
		TaskName:   taskName,
		Decription: description,
		Status:     "todo",
		CreatedAt:  time.Now().Format("2006-01-02"),
		UpdatedAt:  time.Now().Format("2006-01-02"),
	}

	os.Truncate("task.json", 0)      //清空文件
	file.Seek(0, 0)                  //将文件指针移动到开头
	tasks = append(tasks, taskModel) //添加任务

	//再次编码
	err = json.NewEncoder(file).Encode(tasks)
	logger.WarningExitIF("compilation failure", err)

	//格式化
	changejson.Beauty(file)
}
