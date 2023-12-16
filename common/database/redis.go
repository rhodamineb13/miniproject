package database

import "github.com/go-redis/redis/v8"

func NewRedisClient() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})

	return r
}
