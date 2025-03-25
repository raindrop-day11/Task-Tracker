package mark

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	task "task_tracker/app/models/Task"
	"task_tracker/pkg/changejson"
	"time"
)

func MarkInProgress(idstr string) {
	id, _ := strconv.Atoi(idstr)

	//打开文件
	file, err := os.OpenFile("task.json", os.O_RDWR, 0644)
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
	//找到id的任务，更改状态
	var num = 0
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == int64(id) {
			num = 1
			tasks[i].Status = "in-progress"
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02")
			break
		}
	}
	if num == 0 {
		fmt.Println("task does not exist")
		os.Exit(1)
	}
	//清空文件
	os.Truncate("task.json", 0)
	//指针回到开头
	file.Seek(0, 0)
	//编码写入
	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Println("compilation failure")
		os.Exit(1)
	}
	//格式化
	changejson.Beauty(file)
}
