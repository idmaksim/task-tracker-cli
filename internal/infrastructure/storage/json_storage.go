package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/idmaksim/task-tracker-cli/internal/domain/models"
	"github.com/idmaksim/task-tracker-cli/internal/domain/repositories"
)

type jsonTask struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type JSONStorage struct {
	filePath string
	mu       sync.RWMutex
}

func NewJSONStorage(filePath string) *JSONStorage {
	return &JSONStorage{filePath: filePath}
}

func (s *JSONStorage) Init() error {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return s.writeTasks([]*models.Task{})
	}
	return nil
}

func (s *JSONStorage) Create(task *models.Task) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks, err := s.readTasks()
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

	if err := s.writeTasks(tasks); err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (s *JSONStorage) Update(task *models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks, err := s.readTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			return s.writeTasks(tasks)
		}
	}

	return fmt.Errorf("task with id %d not found", task.ID)
}

func (s *JSONStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks, err := s.readTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return s.writeTasks(tasks)
		}
	}

	return fmt.Errorf("task with id %d not found", id)
}

func (s *JSONStorage) FindByID(id int) (*models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks, err := s.readTasks()
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, fmt.Errorf("task with id %d not found", id)
}

func (s *JSONStorage) FindAll(filter repositories.TaskFilter) ([]*models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks, err := s.readTasks()
	if err != nil {
		return nil, err
	}

	if filter.Status == "" {
		return tasks, nil
	}

	var filtered []*models.Task
	for _, task := range tasks {
		if task.Status == filter.Status {
			filtered = append(filtered, task)
		}
	}

	return filtered, nil
}

func (s *JSONStorage) readTasks() ([]*models.Task, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	var jsonTasks []jsonTask
	if err := json.Unmarshal(data, &jsonTasks); err != nil {
		return nil, err
	}

	var tasks []*models.Task
	for _, jt := range jsonTasks {
		task := &models.Task{
			ID:          jt.ID,
			Description: jt.Description,
			Status:      jt.Status,
		}
		if jt.CreatedAt != "" {
			if t, err := time.Parse(time.RFC3339, jt.CreatedAt); err == nil {
				task.CreatedAt = t
			}
		}
		if jt.UpdatedAt != "" {
			if t, err := time.Parse(time.RFC3339, jt.UpdatedAt); err == nil {
				task.UpdatedAt = t
			}
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *JSONStorage) writeTasks(tasks []*models.Task) error {
	var jsonTasks []jsonTask
	for _, task := range tasks {
		jt := jsonTask{
			ID:          task.ID,
			Description: task.Description,
			Status:      task.Status,
		}
		if !task.CreatedAt.IsZero() {
			jt.CreatedAt = task.CreatedAt.Format(time.RFC3339)
		}
		if !task.UpdatedAt.IsZero() {
			jt.UpdatedAt = task.UpdatedAt.Format(time.RFC3339)
		}
		jsonTasks = append(jsonTasks, jt)
	}

	data, err := json.Marshal(jsonTasks)
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}
