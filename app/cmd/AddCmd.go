package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	task "task_tracker/app/models/Task"
	"task_tracker/pkg/changejson"
	"time"
)

func HandleAdd(args []string) {
	taskName := args[0]
	description := args[1]

	var tasks []task.Task
	//打开文件
	file, err := os.OpenFile("task.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	//解码
	dec := json.NewDecoder(file)
	err = dec.Decode(&tasks)
	if err != nil && err.Error() != "EOF" {
		fmt.Println(err)
		os.Exit(1)
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
	enc := json.NewEncoder(file)
	err = enc.Encode(tasks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//格式化
	changejson.Beauty(file)
}
