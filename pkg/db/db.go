package db

import (
    "fmt"

    "developer-allocation-system/pkg/utils"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// NewDatabase initializes and returns a database connection.
func NewDatabase(config utils.ConfigDatabase) (*gorm.DB, error) {
    // Correcting the DSN string construction
    dsn := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        config.Host, config.Port, config.User, config.Password, config.Name, config.SSLMode,
    )

    // Open the connection
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}
