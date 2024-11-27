package repositories

import (
    "developer-allocation-system/pkg/models"
    "gorm.io/gorm"
)

type authRepository struct {
    db *gorm.DB
}

// NewAuthRepository creates a new instance of AuthRepository.
func NewAuthRepository(db *gorm.DB) *authRepository {
    return &authRepository{db: db}
}

// CreateUser creates a new user in the database.
func (r *authRepository) CreateUser(user *models.User) error {
    return r.db.Create(user).Error
}

// GetUserByUsername retrieves a user by their username.
func (r *authRepository) GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    err := r.db.Where("username = ?", username).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
