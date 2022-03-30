package database

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var CTX = context.Background()
var RDB *redis.Client

func InitializeRedisInstance(redisUrl, redisPassword string, redisDatabase int) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPassword,
		DB:       redisDatabase,
	})
	// ping the Redis Client
	_, err := RDB.Ping(CTX).Result()

	if err != nil {
		panic("Could not connect to Redis")
	}
}
