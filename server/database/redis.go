package database

import "github.com/redis/go-redis/v9"

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetRedis() *redis.Client {
	return redisClient
}
