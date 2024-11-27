package utils

import (
    "os"
    "strconv"
)

type Config struct {
    Database   ConfigDatabase
    Cache      ConfigCache
    JWTSecret  string
    ServerPort string
    LogLevel   string
}

type ConfigDatabase struct {
    Host     string
    Port     int
    User     string
    Password string
    Name     string
    SSLMode  string
}

type ConfigCache struct {
    Host string
    Port int
}

func LoadConfig() Config {
    return Config{
        Database: ConfigDatabase{
            Host:     getEnv("DATABASE_HOST", "localhost"),
            Port:     getEnvAsInt("DATABASE_PORT", 5432),
            User:     getEnv("DATABASE_USER", "postgres"),
            Password: getEnv("DATABASE_PASSWORD", "password"),
            Name:     getEnv("DATABASE_NAME", "developer_allocation"),
            SSLMode:  getEnv("DATABASE_SSLMODE", "disable"),
        },
        Cache: ConfigCache{
            Host: getEnv("CACHE_HOST", "localhost"),
            Port: getEnvAsInt("CACHE_PORT", 6379),
        },
        JWTSecret:  getEnv("JWT_SECRET", "your_jwt_secret_here"),
        ServerPort: getEnv("SERVER_PORT", "8080"),
        LogLevel:   getEnv("LOG_LEVEL", "info"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
    if value, exists := os.LookupEnv(key); exists {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }
    return defaultValue
}
