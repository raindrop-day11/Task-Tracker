package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var Tasks []Task

type Task struct {
	Id         int64  `json:"id"`
	TaskName   string `json:"taskname"`
	Decription string `json:"decription"`
	Status     string `json:"status"` //todo,in-progress,done
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func GetByID(idstr string) (Task, error) {
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var tasks []Task

	file, err1 := os.Open("task.json")
	if err1 != nil {
		fmt.Println("json can not open")
		os.Exit(1)
	}

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == int64(id) {
			return tasks[i], nil
		}
	}

	return Task{}, errors.New("not found")
}
