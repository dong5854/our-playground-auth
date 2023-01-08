package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v9"

	"github.com/Team-OurPlayground/our-playground-auth/internal/util/customerror"
)

var redisClient *redis.Client

func RedisInitialize() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_ADDRESS"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal(customerror.Wrap(err, customerror.ErrDBConnection, "failed to ping redis"))
	}
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		log.Fatal(customerror.New(customerror.ErrDBConnection, "redisClient has not been initialized"))
	}
	return redisClient
}
