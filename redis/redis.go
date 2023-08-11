package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	client *redis.Client
	ctx    context.Context
}

func NewStorage(addr, pass string, db int) *Storage {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	return &Storage{ctx: ctx, client: client}
}
