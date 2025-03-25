package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	task "task_tracker/app/models/Task"
	"task_tracker/pkg/changejson"
)

func HandleUpdate(args []string) {
	if len(args) != 3 {
		fmt.Println("args count not right")
		os.Exit(1)
	}
	idstr := args[0]
	id, _ := strconv.Atoi(idstr)
	taskname := args[1]
	description := args[2]

	//打开文件
	var tasks []task.Task
	file, err := os.OpenFile("task.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("file can not open")
		os.Exit(1)
	}

	//解码
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("parsing failure")
		os.Exit(1)
	}

	//更新
	var num = 0
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == int64(id) {
			num = 1
			tasks[i].TaskName = taskname
			tasks[i].Decription = description
			break
		}
	}
	if num == 0 {
		fmt.Println("task does not exist")
		os.Exit(1)
	}

	//清除文件中的内容
	os.Truncate("task.json", 0)
	//将文件指针移动到开头
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
