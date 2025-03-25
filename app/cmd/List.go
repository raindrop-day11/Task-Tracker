package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	task "task_tracker/app/models/Task"
)

func HandleList(args []string) {
	length := len(args)
	if length == 0 {
		List()
	} else if length == 1 {
		if args[0] == "done" || args[0] == "todo" || args[0] == "in-progress" {
			ListByStatus(args[0])
		} else {
			fmt.Println("only done,todo or in-progress")
		}
	} else {
		fmt.Println("wrong number of parameters")
	}
}

func List() {
	//打开文件
	file, err := os.OpenFile("task.json", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("json file can not open")
		os.Exit(1)
	}

	//解码
	var tasks []task.Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("parsing failure")
		os.Exit(1)
	}
	//查询并输出
	for i := 0; i < len(tasks); i++ {
		fmt.Printf("taskname: %s | description: %s\n", tasks[i].TaskName, tasks[i].Decription)
	}
}

func ListByStatus(status string) {
	//打开文件
	file, err := os.OpenFile("task.json", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("json file can not open")
		os.Exit(1)
	}

	//解码
	var tasks []task.Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("parsing failure")
		os.Exit(1)
	}
	//查询并输出
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Status == status {
			fmt.Printf("taskname: %s | description: %s\n", tasks[i].TaskName, tasks[i].Decription)
		}
	}
}
