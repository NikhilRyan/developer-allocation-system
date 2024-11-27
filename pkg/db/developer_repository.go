package db

import (
	"developer-allocation-system/pkg/models"
	"developer-allocation-system/pkg/repositories"
	"gorm.io/gorm"
)

type developerRepository struct {
	db *gorm.DB
}

// NewDeveloperRepository creates a new instance of DeveloperRepository.
func NewDeveloperRepository(db *gorm.DB) repositories.DeveloperRepository {
	return &developerRepository{db: db}
}

func (r *developerRepository) GetAll() ([]models.Developer, error) {
	var developers []models.Developer
	err := r.db.Find(&developers).Error
	return developers, err
}

func (r *developerRepository) GetByID(id int) (*models.Developer, error) {
	var developer models.Developer
	err := r.db.First(&developer, id).Error
	return &developer, err
}

func (r *developerRepository) Create(dev *models.Developer) error {
	return r.db.Create(dev).Error
}

func (r *developerRepository) Update(dev *models.Developer) error {
	return r.db.Save(dev).Error
}

func (r *developerRepository) Delete(id int) error {
	return r.db.Delete(&models.Developer{}, id).Error
}

func (r *developerRepository) UpdateAvailability(id int, availability float64) error {
	return r.db.Model(&models.Developer{}).Where("id = ?", id).Update("availability", availability).Error
}
