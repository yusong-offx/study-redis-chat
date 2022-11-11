package components

import (
	"context"
	"log"

	"github.com/go-redis/redis/v9"
)

var RedisClient *redis.Client

func RedisConnect() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "study-redis-chat-some-redis-1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
}
