package usecases

import (
	"time"

	"github.com/idmaksim/task-tracker-cli/internal/domain/models"
	"github.com/idmaksim/task-tracker-cli/internal/domain/repositories"
)

type TaskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(description string) (int, error) {
	task := models.NewTask(description)
	return s.repo.Create(task)
}

func (s *TaskService) UpdateTaskDescription(id int, description string) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	task.Description = description
	task.UpdatedAt = time.Now()
	return s.repo.Update(task)
}

func (s *TaskService) UpdateTaskStatus(id int, status string) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	task.Status = status
	task.UpdatedAt = time.Now()
	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}

func (s *TaskService) GetTasks(status string) ([]*models.Task, error) {
	return s.repo.FindAll(repositories.TaskFilter{Status: status})
}
