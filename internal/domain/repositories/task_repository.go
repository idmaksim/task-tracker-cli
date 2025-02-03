package repositories

import "github.com/idmaksim/task-tracker-cli/internal/domain/models"

type TaskRepository interface {
	Create(task *models.Task) (int, error)
	Update(task *models.Task) error
	Delete(id int) error
	FindByID(id int) (*models.Task, error)
	FindAll(filter TaskFilter) ([]*models.Task, error)
}

type TaskFilter struct {
	Status string
}
