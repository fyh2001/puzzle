package database

import (
	"puzzle/config"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.Settings.Redis.Host + ":" + config.Settings.Redis.Port,
		Password: config.Settings.Redis.Password, // no password set
		DB:       config.Settings.Redis.Db,       // use default DB
	})

}

func GetRedis() *redis.Client {
	return redisClient
}
