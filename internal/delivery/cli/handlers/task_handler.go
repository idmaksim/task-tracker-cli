package handlers

import (
	"fmt"

	"github.com/idmaksim/task-tracker-cli/internal/usecases"
)

type TaskHandler struct {
	service *usecases.TaskService
}

func NewTaskHandler(service *usecases.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(description string) error {
	id, err := h.service.CreateTask(description)
	if err != nil {
		return err
	}

	fmt.Printf("Task added successfully with ID: %d\n", id)
	return nil
}

func (h *TaskHandler) UpdateTaskDescription(id int, description string) error {
	if err := h.service.UpdateTaskDescription(id, description); err != nil {
		return err
	}

	fmt.Println("Task updated successfully")
	return nil
}

func (h *TaskHandler) UpdateTaskStatus(id int, status string) error {
	if err := h.service.UpdateTaskStatus(id, status); err != nil {
		return err
	}

	fmt.Printf("Task marked as %s\n", status)
	return nil
}

func (h *TaskHandler) DeleteTask(id int) error {
	if err := h.service.DeleteTask(id); err != nil {
		return err
	}

	fmt.Println("Task deleted successfully")
	return nil
}

func (h *TaskHandler) ListTasks(status string) error {
	tasks, err := h.service.GetTasks(status)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	fmt.Println("ID\t\tStatus\t\tDescription")
	for _, task := range tasks {
		fmt.Printf("%d\t\t%s\t\t%s\n", task.ID, task.Status, task.Description)
	}

	return nil
}
