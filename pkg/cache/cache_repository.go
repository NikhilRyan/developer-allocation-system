package cache

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"developer-allocation-system/pkg/repositories"
)

type cacheRepository struct {
	client  *redis.Client
	context context.Context
}

// NewCacheRepository creates a new instance of CacheRepository.
func NewCacheRepository(client *redis.Client) repositories.CacheRepository {
	return &cacheRepository{
		client:  client,
		context: context.Background(),
	}
}

func (r *cacheRepository) Get(key string, dest interface{}) error {
	val, err := r.client.Get(r.context, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

func (r *cacheRepository) Set(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(r.context, key, data, 0).Err()
}

func (r *cacheRepository) Delete(key string) error {
	return r.client.Del(r.context, key).Err()
}
