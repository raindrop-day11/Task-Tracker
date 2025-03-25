package cmd

import (
	"encoding/json"
	"os"
	"strconv"
	task "task_tracker/app/models/Task"
	"task_tracker/pkg/changejson"
	"task_tracker/pkg/logger"
	"time"
)

func HandleMark(args []string) {
	if len(args) != 2 {
		logger.WarningExit("only done,todo or in-progress")
	}

	if args[0] == "in-progress" || args[0] == "done" {
		Mark(args)
	} else {
		logger.WarningExit("only done,todo or in-progress")
	}
}

func Mark(args []string) {
	status := args[0]
	idstr := args[1]
	id, _ := strconv.Atoi(idstr)

	//打开文件
	file, err := os.OpenFile("task.json", os.O_RDWR, 0644)
	logger.WarningExitIF("json file can not open", err)

	//解码
	var tasks []task.Task
	err = json.NewDecoder(file).Decode(&tasks)
	logger.WarningExitIF("parsing failure", err)

	//找到id的任务，更改状态
	var num = 0
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == int64(id) {
			num = 1
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02")
			break
		}
	}
	if num == 0 {
		logger.WarningExit("task does not exist")
	}

	//清空文件
	os.Truncate("task.json", 0)
	//指针回到开头
	file.Seek(0, 0)

	//编码写入
	err = json.NewEncoder(file).Encode(tasks)
	logger.WarningExitIF("compilation failure", err)

	//格式化
	changejson.Beauty(file)
}
