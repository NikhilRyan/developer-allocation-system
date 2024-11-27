package repositories

import "developer-allocation-system/pkg/models"

type DeveloperRepository interface {
    GetAll() ([]models.Developer, error)
    GetByID(id int) (*models.Developer, error)
    Create(dev *models.Developer) error
    Update(dev *models.Developer) error
    Delete(id int) error
    UpdateAvailability(id int, availability float64) error
}
