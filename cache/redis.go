package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{Client: client}
}
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	value, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("cache miss for key: %s", key)
	} else if err != nil {
		return "", fmt.Errorf("error getting key %s: %v", key, err)
	}
	return value, nil
}

func (r *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := r.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("error setting key %s in Redis: %v", key, err)
	}
	return nil
}
