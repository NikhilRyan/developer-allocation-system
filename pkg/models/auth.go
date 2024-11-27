package models

import (
    "time"
)

// User represents a user in the system for authentication.
type User struct {
    ID        int       `json:"id" gorm:"primaryKey"`
    Username  string    `json:"username" binding:"required" gorm:"unique"`
    Password  string    `json:"password" binding:"required"`
    Role      string    `json:"role"` // e.g., admin, user
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// Credentials represents user login credentials.
type Credentials struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}
