package repositories

import (
	"developer-allocation-system/pkg/models"
	"gorm.io/gorm"
)

// UserRepository defines the interface for user-related operations.
type UserRepository interface {
	Create(user *models.User) error
	GetByUsername(username string) (*models.User, error)
}

// userRepository implements the UserRepository interface.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create adds a new user to the database.
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetByUsername retrieves a user by their username.
func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
