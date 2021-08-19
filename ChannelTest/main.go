package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go producer(ch)
	go consumption(ch)
	time.Sleep(2000)

}

func consumption(c chan int) {
	select {
	case <-c:
		fmt.Println(<-c)
	}
}

func producer(c chan int) {
	defer fmt.Println("生产任务结束")
	for i := 1; i <= 10; i++ {
		fmt.Println("生产任务:", i)
		c <- i
	}
	close(c)
}
