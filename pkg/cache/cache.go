package cache

import (
    "context"
    "fmt"

    "github.com/go-redis/redis/v8"
    "developer-allocation-system/pkg/utils"
)

// NewCache initializes and returns a Redis client.
func NewCache(config utils.ConfigCache) (*redis.Client, error) {
    client := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
        // Password: config.Password, // no password set
        // DB:       config.DB,       // use default DB
    })

    // Test the connection
    _, err := client.Ping(context.Background()).Result()
    if err != nil {
        return nil, err
    }

    return client, nil
}
