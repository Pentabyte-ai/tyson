package database

import (
	"context"
	"fmt"
	"log"
	"github.com/go-redis/redis/v8"
)

//connecting to RedisClient
func RedisClient()*redis.Client{
    redisHost := "brewdctf7h5b6avp6dnm-redis.services.clever-cloud.com"
    redisPort := "3630"
    redisPassword := "6VIqwXR5je4NdGqCMJq"

    // Construct the Redis client options
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
        Password: redisPassword,
        DB:       0, // use default DB
    })

    // Test the connection
    ctx := context.Background()
    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
	return rdb
}