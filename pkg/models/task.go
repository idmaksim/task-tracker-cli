package models

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func NewTask(description string) Task {
	return Task{
		Description: description,
		Status:      "todo",
	}
}
