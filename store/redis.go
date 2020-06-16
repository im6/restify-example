package store

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

// RedisDB is the cache connection instance
var RedisDB *redis.Client

// ConnectRedis init conn
func ConnectRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	RedisDB = rdb
}
