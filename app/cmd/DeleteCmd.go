package cmd

import (
	"encoding/json"
	"os"
	"strconv"
	task "task_tracker/app/models/Task"
	"task_tracker/pkg/changejson"
	"task_tracker/pkg/logger"
)

func HandleDelete(args []string) {
	if len(args) != 1 {
		logger.WarningExit("wrong number of parameters")
	}
	idstr := args[0]
	id, _ := strconv.Atoi(idstr)

	//打开文件
	file, err := os.OpenFile("task.json", os.O_RDWR, 0644)
	logger.WarningExitIF("json file can not open", err)
	defer file.Close()

	//解析
	var tasks []task.Task
	err = json.NewDecoder(file).Decode(&tasks)
	logger.WarningExitIF("parsing failure", err)

	//删除
	var num = 0
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == int64(id) {
			num = 1
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	if num == 0 {
		logger.WarningExit("task does not exist")
	}

	//清空内容
	os.Truncate("task.json", 0)
	//指针移动到开头
	file.Seek(0, 0)

	//编码
	err = json.NewEncoder(file).Encode(tasks)
	logger.WarningExitIF("compilation failure", err)

	//格式化
	changejson.Beauty(file)
}
