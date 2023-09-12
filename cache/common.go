package cache

import (
	"context"
	"mail/config"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func init() {
	//db, _ := strconv.ParseUint(config.Redis.DBName, 10, 64)
	cliet := redis.NewClient(&redis.Options{
		Addr: config.Redis.Address,
		// password
		//DB: int(db),
	})
	_, err := cliet.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	RedisClient = cliet
}
