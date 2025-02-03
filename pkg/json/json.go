package json

import (
	"encoding/json"
	"os"
	"task-tracker-cli/pkg/models"
	"task-tracker-cli/pkg/types"
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

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
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
