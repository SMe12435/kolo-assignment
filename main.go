package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "io/ioutil"
	_ "kolo-assignment/config"
	"kolo-assignment/repl"
	"kolo-assignment/server"
)

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "valuefjsdjsnd", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func main() {
	//using routine to initialise the API server. Used this approach so that
	go server.Init()
	//ExampleClient()
	//initializes the REPL
	repl.Init()

}
