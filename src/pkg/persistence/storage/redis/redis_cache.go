package cache

import (
	"context"
	"encoding/json"
	"genesis_test_case/src/pkg/usecase"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	expires time.Duration
	client  *redis.Client
}

func NewRedisCache(host string, db int, exp time.Duration) usecase.Cache {
	return &redisCache{
		client:  getRedisClient(host, db),
		expires: exp,
	}
}

func getRedisClient(host string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: host,
		DB:   db,
	})
}

func (r *redisCache) GetCache(key string) ([]byte, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(val), nil
}

func (r *redisCache) SetCache(key string, value interface{}) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	r.client.Set(context.Background(), key, json, r.expires)

	return nil
}
