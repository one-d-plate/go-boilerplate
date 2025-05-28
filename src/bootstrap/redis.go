package bootstrap

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type redisClient struct{}

func NewRedisClient() *redisClient {
	return &redisClient{}
}

func (r *redisClient) Connect() (*redis.Client, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		cancel()
		log.Printf("❌ Failed to connect to Redis: %v", err)
		return nil, err
	}

	log.Println("✅ Connected to Redis with auth")
	return client, nil
}
