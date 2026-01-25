package cache

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		log.Fatal("REDIS_ADDR not set")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Redis")
	return rdb
}
