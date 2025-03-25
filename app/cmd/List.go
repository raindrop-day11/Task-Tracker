package cmd

import (
	"encoding/json"
	"os"
	task "task_tracker/app/models/Task"
	"task_tracker/pkg/logger"

	"github.com/fatih/color"
)

func HandleList(args []string) {
	length := len(args)
	if length == 0 {
		List()
	} else if length == 1 {
		if args[0] == "done" || args[0] == "todo" || args[0] == "in-progress" {
			ListByStatus(args[0])
		} else {
			logger.WarningExit("only done,todo or in-progress")
		}
	} else {
		logger.WarningExit("wrong number of parameters")
	}
}

func List() {
	//打开文件
	file, err := os.OpenFile("task.json", os.O_RDONLY, 0644)
	logger.WarningExitIF("json file can not open", err)

	//解码
	var tasks []task.Task
	err = json.NewDecoder(file).Decode(&tasks)
	logger.WarningExitIF("parsing failure", err)

	//查询并输出
	color.New(color.BgHiGreen).Fprintln(os.Stdout, "tasks:")
	for i := 0; i < len(tasks); i++ {
		color.New(color.BgHiGreen).Fprintf(os.Stdout, "taskname: %s | description: %s\n", tasks[i].TaskName, tasks[i].Decription)
	}
}

func ListByStatus(status string) {
	//打开文件
	file, err := os.OpenFile("task.json", os.O_RDONLY, 0644)
	logger.WarningExitIF("json file can not open", err)

	//解码
	var tasks []task.Task
	err = json.NewDecoder(file).Decode(&tasks)
	logger.WarningExitIF("parsing failure", err)

	//查询并输出
	color.New(color.BgHiGreen).Fprintln(os.Stdout, "tasks:")
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Status == status {
			color.New(color.BgHiGreen).Fprintf(os.Stdout, "taskname: %s | description: %s\n", tasks[i].TaskName, tasks[i].Decription)
		}
	}
}
