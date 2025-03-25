package task

import (
	"encoding/json"
	"fmt"
	"os"
	"task_tracker/pkg/changejson"
)

func (task Task) AddToJSON() {
	file, err := os.OpenFile("task.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//解码
	dec := json.NewDecoder(file)
	err = dec.Decode(&Tasks)
	if err != nil && err.Error() != "EOF" {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Truncate("task.json", 0) //清空文件
	file.Seek(0, 0)             //将文件指针移动到开头
	Tasks = append(Tasks, task) //添加任务

	//再次编码
	enc := json.NewEncoder(file)
	err = enc.Encode(Tasks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//格式化
	changejson.Beauty(file)
}
