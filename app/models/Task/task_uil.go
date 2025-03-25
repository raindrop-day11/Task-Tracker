package task

import (
	"encoding/json"
	"os"
	"task_tracker/pkg/changejson"
)

func (task Task) WriteToJSON() error {
	file, err := os.OpenFile("task.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	//解码
	dec := json.NewDecoder(file)
	err = dec.Decode(&Tasks)
	if err != nil && err.Error() != "EOF" {
		return err
	}

	os.Truncate("task.json", 0) //清空文件
	file.Seek(0, 0)             //将文件指针移动到开头
	Tasks = append(Tasks, task) //添加任务

	//再次编码
	enc := json.NewEncoder(file)
	err = enc.Encode(Tasks)
	if err != nil {
		return err
	}
	changejson.Beauty()
	return nil
}
