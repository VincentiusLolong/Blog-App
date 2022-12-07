package configs

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)

func contectx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	a, b := contectx()
	defer b()
	_, err := rdb.Ping(a).Result()
	if err != nil {
		log.Fatal(err)
	}
	return rdb
}

func RedisGet(id string) (string, error) {
	a, b := contectx()
	defer b()
	val, err := ConnectRedis().Get(a, id).Result()
	if err != nil {
		log.Fatal(err)
	}
	return val, nil
}

func RedisSet(id, token string, ttl time.Duration) error {
	a, b := contectx()
	defer b()
	if err != nil {
		log.Fatal(err)
	}
	err := ConnectRedis().Set(a, id, token, ttl).Err()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
