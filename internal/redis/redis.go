package redis

import (
	"context"
	"encoding/json"
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

func (r *RedisService) SetJSON(key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.Client.Set(r.Ctx, key, data, ttl).Err()
}

func (r *RedisService) GetJSON(key string, dest any) (bool, error) {
	val, err := r.Client.Get(r.Ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(val), dest)
	return true, err
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
