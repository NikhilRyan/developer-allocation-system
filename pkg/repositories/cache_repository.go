package repositories

type CacheRepository interface {
    Get(key string, dest interface{}) error
    Set(key string, value interface{}) error
    Delete(key string) error
}
