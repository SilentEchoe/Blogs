package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.Get(ctx, "devops_dashboard_build").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	str1 := []byte(`{"total":"45","success":22}`)

	var revMsg1 test
	err = json.Unmarshal(str1, &revMsg1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("test")
	fmt.Println(revMsg1)

	str := []byte(`{"FromUid1":"100","FromName":"HANASHAN"}`)

	var revMsg receiveMessage
	err = json.Unmarshal(str, &revMsg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(revMsg)

}

//接收普通消息结构体
type receiveMessage struct {
	FromUid1 string //发送者uid
	FromName string //发送者名字
}

type test struct {
	Total   string
	success int
}

type Data struct {
	date    string
	success int
	failure int
	total   int
}
