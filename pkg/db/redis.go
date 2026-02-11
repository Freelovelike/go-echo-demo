package db

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis(addr string) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "hwc20010616", // no password set
		DB:       0,             // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Redis.Ping(ctx).Err(); err != nil {
		panic("Redis 连接失败: " + err.Error())
	}
	println("Redis 连接成功")
}
