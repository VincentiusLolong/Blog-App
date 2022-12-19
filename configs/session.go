package configs

import (
	"context"
	"log"
	"time"

	redisClient "github.com/go-redis/redis/v9"
	// redisStore "github.com/gofiber/storage/redis"
)

func contectx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}

//	func Redisstore() *redisStore.Storage {
//		client := redisStore.New(redisStore.Config{
//			Host: "127.0.0.1:6379",
//		})
//		if client != nil {
//			log.Fatal(err)
//		}
//		return client
//	}
func ConnectRedis() *redisClient.Client {
	rdb := redisClient.NewClient(&redisClient.Options{
		Addr:     "127.0.0.1:6379",
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

func RedisSet(id, token string) error {
	a, b := contectx()
	defer b()
	if err != nil {
		log.Fatal(err)
	}
	err := ConnectRedis().Set(a, id, token, 30*24*time.Hour).Err()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func RedisDelete(id string) error {
	a, b := contectx()
	defer b()
	err := ConnectRedis().Del(a, id).Err()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
