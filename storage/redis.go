package storage

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type URLModel struct {
	ID    string
	Value string
}

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(addr, pass string, db int) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	return &RedisStorage{client: client}
}

func (rs *RedisStorage) Get(ctx context.Context, id string) (*URLModel, error) {
	return nil, nil
}

func (rs *RedisStorage) Set(ctx context.Context, url string) error {
	return nil
}
