package cache

import (
	"context"
	"github.com/TskFok/GinApi/app/global"
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.RedisHost,
		Password: global.RedisPassword,
		DB:       10,
	})

	ctx := context.Background()

	cmd := client.Ping(ctx)

	if nil != cmd.Err() {
		panic(cmd.Err())
	}

	return client
}
