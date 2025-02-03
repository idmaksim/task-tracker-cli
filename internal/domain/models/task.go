package models

import (
	"time"

	"github.com/idmaksim/task-tracker-cli/pkg/constants"

)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

func NewTask(description string) *Task {
	now := time.Now()
	return &Task{
		Description: description,
		Status:      constants.Todo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

}
