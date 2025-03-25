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

func HandleUpdate(args []string) {
	if len(args) != 3 {
		logger.WarningExit("wrong number of parameters")
	}
	idstr := args[0]
	id, _ := strconv.Atoi(idstr)
	taskname := args[1]
	description := args[2]

	//打开文件
	var tasks []task.Task
	file, err := os.OpenFile("task.json", os.O_RDWR, 0644)
	logger.WarningExitIF("json file can not open", err)
	defer file.Close()

	//解码
	err = json.NewDecoder(file).Decode(&tasks)
	logger.WarningExitIF("parsing failure", err)

	//更新
	var num = 0
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == int64(id) {
			num = 1
			tasks[i].TaskName = taskname
			tasks[i].Decription = description
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02")
			break
		}
	}
	if num == 0 {
		logger.WarningExit("task does not exist")
	}

	//清除文件中的内容
	os.Truncate("task.json", 0)
	//将文件指针移动到开头
	file.Seek(0, 0)

	//编码
	err = json.NewEncoder(file).Encode(tasks)
	logger.WarningExitIF("compilation failure", err)

	//格式化
	changejson.Beauty(file)
}
