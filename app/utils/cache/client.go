package cache

import (
	"context"
	"fmt"
	"github.com/TskFok/GinApi/app/utils/conf"
	"github.com/redis/go-redis/v9"
	"time"
)

func getClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisHost,
		Password: conf.RedisPassword,
		DB:       10,
	})

	return client
}

func Get(key string) string {
	client := getClient()
	ctx := context.Background()
	result, err := client.Get(ctx, key).Result()

	if nil != err {
		fmt.Println(err)
	}

	return result
}

func Set(key string, value string) {
	client := getClient()
	ctx := context.Background()
	err := client.Set(ctx, key, value, 0).Err()

	if nil != err {
		fmt.Println(err)
	}
}

func SetNx(key string, value string, limit int64) bool {
	client := getClient()
	ctx := context.Background()

	limitTime := time.Duration(limit) * time.Second

	set, err := client.SetNX(ctx, key, value, limitTime).Result()

	if nil != err {
		fmt.Println(err)
	}
	return set
}
