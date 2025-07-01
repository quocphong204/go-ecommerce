package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var RedisCtx = context.Background()

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379", // dùng tên service Redis trong docker-compose
	})
	_, err := RedisClient.Ping(RedisCtx).Result()
	if err != nil {
		panic(fmt.Sprintf("❌ Failed to connect to Redis: %v", err))
	}
	fmt.Println("✅ Connected to Redis!")
}
