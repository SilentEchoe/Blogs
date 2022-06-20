package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"net/http"
)

var ctx = context.Background()

func ExampleClient(wr http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis.default:6379",
		//Addr:     "10.3.70.149:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
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

	wr.Write([]byte("hello"))
	// Output: key value
	// key2 does not exist
}

func main() {
	http.HandleFunc("/", ExampleClient)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
