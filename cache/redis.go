package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

var ctx = context.Background()

type Params struct {
	Host string
	Port string
	Db   int
}

func GetRedis(params Params) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     params.Host + params.Port,
		Password: "",        // no password set
		DB:       params.Db, // use default DB
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
