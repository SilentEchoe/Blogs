package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {

	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:3309",
		DB:   0, // use default DB
	})

	err := client.LPush(context.Background(), "OperationLog", "admin访问Devops").Err()
	if err != nil {
		fmt.Println(err)
	}
}
