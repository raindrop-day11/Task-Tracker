package task

type Task struct {
	Id         int64  `json:"id"`
	TaskName   string `json:"taskname"`
	Decription string `json:"decription"`
	Status     string `json:"status"` //todo,in-progress,done
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
