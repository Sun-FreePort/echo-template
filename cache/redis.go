package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

var ctx = context.Background()

type Parameters struct {
	db int `default:"0"`
}

func GetRedis(params *Parameters) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",        // no password set
		DB:       params.db, // use default DB
	})

	return rdb
}

func Get(rdb *redis.Client, key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		panic(err)
	}

	return val
}

func Set(rdb *redis.Client, key string, val string) string {
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		panic(err)
	}

	return val
}

func SetExpiration(rdb *redis.Client, key string, val string, expiration time.Duration) string {
	err := rdb.Set(ctx, key, val, expiration).Err()
	if err != nil {
		panic(err)
	}

	return val
}

func Expire(rdb *redis.Client, key string, expiration time.Duration) string {
	err := rdb.Expire(ctx, key, expiration).Err()
	if err != nil {
		panic(err)
	}

	return key
}

func Delete(rdb *redis.Client, key string) bool {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}
	return true
}
