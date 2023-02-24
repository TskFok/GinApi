package cache

import (
	"context"
	"github.com/TskFok/GinApi/app/utils/conf"
	"github.com/redis/go-redis/v9"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     conf.RedisHost,
		Password: conf.RedisPassword,
		DB:       10,
	})
}

func Get(key string) string {
	ctx := context.Background()
	result, err := client.Get(ctx, key).Result()

	if nil != err {
		panic(err)
	}

	return result
}

func Set(key string, value string) {
	ctx := context.Background()
	err := client.Set(ctx, key, value, 0).Err()

	if nil != err {
		panic(err)
	}
}

func SetNx(key string, value string, limit int64) bool {
	ctx := context.Background()

	limitTime := time.Duration(limit) * time.Second

	set, err := client.SetNX(ctx, key, value, limitTime).Result()

	if nil != err {
		panic(err)
	}
	return set
}
