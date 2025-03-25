package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	task "task_tracker/app/models/Task"
	"task_tracker/pkg/changejson"
)

func HandleDelete(args []string) {
	if len(args) != 1 {
		fmt.Println("wrong number of parameters")
		os.Exit(1)
	}
	idstr := args[0]
	id, _ := strconv.Atoi(idstr)

	//打开文件
	file, err := os.OpenFile("task.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("json can not open")
		os.Exit(1)
	}
	defer file.Close()

	//解析
	var tasks []task.Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("parsing failure")
		os.Exit(1)
	}

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
		fmt.Println("task does not exist")
		os.Exit(1)
	}

	//清空内容
	os.Truncate("task.json", 0)
	//指针移动到开头
	file.Seek(0, 0)

	//编码
	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Println("compilation failure")
		os.Exit(1)
	}

	//格式化
	changejson.Beauty(file)
}
