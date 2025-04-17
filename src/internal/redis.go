package internal

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

// NewRedisClient initializes and returns a RedisClient
func NewRedisClient() *RedisClient {
	addr := os.Getenv("REDIS_ADDRESS")
	pass := os.Getenv("REDIS_PASSWORD")
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		log.Fatalf("Invalid REDIS_DB value: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	return &RedisClient{
		Client: rdb,
		Ctx:    context.Background(),
	}
}

func (r *RedisClient) Initialize() {
	r.Client.Set(r.Ctx, "foo", "bar", 0)
	r.Client.Set(r.Ctx, "key", "value", 0)
}

// GetClient exposes the Redis client for routes to use
func (r *RedisClient) GetClient() *redis.Client {
	return r.Client
}
