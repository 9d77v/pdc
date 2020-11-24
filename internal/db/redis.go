package db

import redis "github.com/go-redis/redis/v8"

func initRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0, // use default DB
	})
}
