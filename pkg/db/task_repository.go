package db

import (
	"developer-allocation-system/pkg/models"
	"developer-allocation-system/pkg/repositories"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository creates a new instance of TaskRepository.
func NewTaskRepository(db *gorm.DB) repositories.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetByID(id int) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *taskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id int) error {
	return r.db.Delete(&models.Task{}, id).Error
}
