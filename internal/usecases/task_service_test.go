package usecases

import (
	"fmt"
	"testing"

	"github.com/idmaksim/task-tracker-cli/internal/domain/models"
	"github.com/idmaksim/task-tracker-cli/internal/domain/repositories"
	"github.com/idmaksim/task-tracker-cli/pkg/constants"
)

type mockTaskRepository struct {
	tasks  map[int]*models.Task
	lastID int
}

func newMockTaskRepository() *mockTaskRepository {
	return &mockTaskRepository{
		tasks:  make(map[int]*models.Task),
		lastID: 0,
	}
}

func (m *mockTaskRepository) Create(task *models.Task) (int, error) {
	m.lastID++
	task.ID = m.lastID
	m.tasks[task.ID] = task
	return task.ID, nil
}

func (m *mockTaskRepository) Update(task *models.Task) error {
	if _, exists := m.tasks[task.ID]; !exists {
		return fmt.Errorf("task not found")
	}
	m.tasks[task.ID] = task
	return nil
}

func (m *mockTaskRepository) Delete(id int) error {
	if _, exists := m.tasks[id]; !exists {
		return fmt.Errorf("task not found")
	}
	delete(m.tasks, id)
	return nil
}

func (m *mockTaskRepository) FindByID(id int) (*models.Task, error) {
	task, exists := m.tasks[id]
	if !exists {
		return nil, fmt.Errorf("task not found")
	}
	return task, nil
}

func (m *mockTaskRepository) FindAll(filter repositories.TaskFilter) ([]*models.Task, error) {
	var result []*models.Task
	for _, task := range m.tasks {
		if filter.Status == "" || task.Status == filter.Status {
			result = append(result, task)
		}
	}
	return result, nil
}

func TestCreateTask(t *testing.T) {
	repo := newMockTaskRepository()
	service := NewTaskService(repo)

	description := "Test task"
	id, err := service.CreateTask(description)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if id != 1 {
		t.Errorf("Expected ID 1, got %d", id)
	}

	task, _ := repo.FindByID(id)
	if task.Description != description {
		t.Errorf("Expected description %s, got %s", description, task.Description)
	}
	if task.Status != constants.Todo {
		t.Errorf("Expected status %s, got %s", constants.Todo, task.Status)
	}
}

func TestUpdateTaskDescription(t *testing.T) {
	repo := newMockTaskRepository()
	service := NewTaskService(repo)

	id, _ := service.CreateTask("Initial description")
	newDescription := "Updated description"

	err := service.UpdateTaskDescription(id, newDescription)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	task, _ := repo.FindByID(id)
	if task.Description != newDescription {
		t.Errorf("Expected description %s, got %s", newDescription, task.Description)
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	repo := newMockTaskRepository()
	service := NewTaskService(repo)

	id, _ := service.CreateTask("Test task")
	newStatus := constants.InProgress

	err := service.UpdateTaskStatus(id, newStatus)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	task, _ := repo.FindByID(id)
	if task.Status != newStatus {
		t.Errorf("Expected status %s, got %s", newStatus, task.Status)
	}
}

func TestDeleteTask(t *testing.T) {
	repo := newMockTaskRepository()
	service := NewTaskService(repo)

	id, _ := service.CreateTask("Test task")

	err := service.DeleteTask(id)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = repo.FindByID(id)
	if err == nil {
		t.Error("Expected error when finding deleted task")
	}
}

func TestGetTasks(t *testing.T) {
	repo := newMockTaskRepository()
	service := NewTaskService(repo)

	service.CreateTask("Task 1")
	id2, _ := service.CreateTask("Task 2")
	service.UpdateTaskStatus(id2, constants.InProgress)

	tests := []struct {
		name          string
		status        string
		expectedCount int
	}{
		{"All tasks", "", 2},
		{"Todo tasks", constants.Todo, 1},
		{"In progress tasks", constants.InProgress, 1},
		{"Done tasks", constants.Done, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tasks, err := service.GetTasks(tt.status)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if len(tasks) != tt.expectedCount {
				t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(tasks))
			}
		})
	}
}
