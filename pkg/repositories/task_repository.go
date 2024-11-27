package repositories

import "developer-allocation-system/pkg/models"

type TaskRepository interface {
    GetAll() ([]models.Task, error)
    GetByID(id int) (*models.Task, error)
    Create(task *models.Task) error
    Update(task *models.Task) error
    Delete(id int) error
}
