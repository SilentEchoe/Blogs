package main

import (
	"context"
	"fmt"
	"time"
)

// 如果 chan 中还有数据,那么从这个 chan 接收数据的时候就不会阻塞

// nil 是 chan 的零值,是一种特殊的 chan, 对值是 nil 的 chan 的发送接受调用者总是会阻塞
func main() {

	// chan为nil 的时候,从 nil chan 中接受（读取,获取）数据时，调用者会被永远阻塞
	// 如果 chan 已经被 close,并且队列中没有缓存的元素,那么返回 true,false
	var ch = make(chan int, 10)
	done := make(chan bool)
	go producers(ch, done)
	go consumers(ch, done)
	<-done
	//time.Sleep(10 * time.Second)

	// 借助 context.WithCancel 上下文
	date := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan bool)

	go SearchTarget(ctx, date, 7, resultChan)
	go SearchTarget(ctx, date, 9, resultChan)
	select {
	case <-resultChan:
		fmt.Println("Find item")
		// 通知上下文
		cancel()
	}

}

func producers(ch chan int, done chan bool) {
	for i := 0; i < 100; i++ {
		fmt.Println("生产", i)
		ch <- i
	}
	done <- true
}

func consumers(ch chan int, done chan bool) {
	for {
		select {
		case r := <-ch:
			if r > 0 {
				fmt.Println("消费", r)
			}
		case <-done:
			return
		}
	}

}

func SearchTarget(ctx context.Context, data []int, target int, resultChan chan bool) {
	for _, v := range data {

		// 监听上下文是否取消任务
		select {
		case <-ctx.Done():
			return
		default:
		}
		time.Sleep(1 * time.Second)
		if v == target {
			resultChan <- true
			return
		}
	}
}
