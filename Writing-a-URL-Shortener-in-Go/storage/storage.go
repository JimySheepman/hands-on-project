package storage

import "github.com/go-redis/redis"

type Storage interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
	Exists(key string) (bool, error)
}

type RedisStorage struct {
	client *redis.Client
}

func (r *RedisStorage) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *RedisStorage) Set(key string, value interface{}) error {
	return r.client.Set(key, value, 0).Err()
}

func (r *RedisStorage) Exists(key string) (bool, error) {
	exists, err := r.client.Exists(key).Result()
	if err != nil {
		return false, err
	}

	return exists == 1, nil
}

func (r RedisStorage) String() string {
	return "RedisStorage"
}

func New(address string) *RedisStorage {
	var storage = &RedisStorage{}
	storage.client = redis.NewClient(&redis.Options{
		Addr: address,
	})

	return storage
}
