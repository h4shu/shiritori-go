package db

import (
	"context"
	"os"

	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
}

func NewRedisStore(ctx context.Context, rdb *redis.Client) (*redisstore.RedisStore, error) {
	return redisstore.NewRedisStore(ctx, rdb)
}
