package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

var client *redis.Client
var once sync.Once

func Init() {
	once.Do(func() {
		var ctx = context.Background()
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		err := rdb.Set(ctx, "key", "valuefjsdjsnd", 0).Err()
		if err != nil {
			panic(err)
		}

	})
}

func Client() *redis.Client {
	return client
}
