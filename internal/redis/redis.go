package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisService(addr string) *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisService{
		Client: rdb,
		Ctx:    context.Background(),
	}
}

func (r *RedisService) Set(key string, value any, ttl time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, ttl).Err()
}

func (r *RedisService) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

func (r *RedisService) Delete(key string) error {
	return r.Client.Del(r.Ctx, key).Err()
}

func (r *RedisService) Exists(key string) (bool, error) {
	count, err := r.Client.Exists(r.Ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
