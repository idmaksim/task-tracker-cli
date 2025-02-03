package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/idmaksim/task-tracker-cli/pkg/models"
	"github.com/idmaksim/task-tracker-cli/pkg/types"
)

var dataPath = "data.json"

func StartUp() {
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		writeTasks([]models.Task{})
	}
}

func readTasks() ([]models.Task, error) {
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	if err = json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func writeTasks(tasks []models.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	if err = os.WriteFile(dataPath, data, 0644); err != nil {
		return err
	}
	return nil
}

func ReadAllData(options types.FindAllOptions) ([]models.Task, error) {
	tasks, err := readTasks()
	if err != nil {
		return nil, err
	}
	return filterByStatus(tasks, options.Status), nil

}

func WriteData(task models.Task) (int, error) {
	tasks, err := readTasks()
	if err != nil {
		return 0, err
	}

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	task.ID = maxID + 1
	tasks = append(tasks, task)
	err = writeTasks(tasks)
	if err != nil {
		return 0, err
	}
	return task.ID, nil
}

func DeleteData(id int) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	var deleted models.Task

	for i, task := range tasks {
		if task.ID == id {
			deleted = task
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	if deleted.ID == 0 {
		return fmt.Errorf("task with id %d not found", id)
	}

	return writeTasks(tasks)
}

func filterByStatus(tasks []models.Task, status string) []models.Task {
	if status == "" {
		return tasks
	}
	var filtered []models.Task
	for _, task := range tasks {
		if task.Status == status {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func UpdateTaskStatus(id int, status string) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	taskFound := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			taskFound = true
			break
		}
	}

	if !taskFound {
		return fmt.Errorf("задача с id %d не найдена", id)
	}

	return writeTasks(tasks)
}

func UpdateTaskDescription(id int, description string) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	taskFound := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			taskFound = true
			break
		}
	}

	if !taskFound {
		return fmt.Errorf("task with id %d not found", id)
	}

	return writeTasks(tasks)
}
