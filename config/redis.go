package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var RedisContext = context.Background()

func InitRedis() {
	redisDb, osErr := strconv.Atoi(os.Getenv("redis_db"))
	if osErr != nil {
		redisDb = 0
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_host") + ":" + os.Getenv("redis_port"), // Redis server address
		Password: os.Getenv("redis_password"),                             // Redis password (leave empty if none)
		DB:       redisDb,                                                 // Default DB
	})

	_, err := RedisClient.Ping(RedisContext).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Successfully connected to Redis!")
}
