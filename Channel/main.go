package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)

	//var wg sync.WaitGroup
	//wg.Add(2)
	go producer(ch)
	//go consume(ch)

	for c := range ch {
		fmt.Println(c)
	}

	time.Sleep(15000)

	//wg.Wait()
}

func producer(c chan int) {
	defer fmt.Println("生产任务结束")
	for i := 1; i <= 10; i++ {
		fmt.Println("生产任务:", i)
		c <- i
	}
	close(c)
}

func consume(c chan int) {
	for {
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("消费任务 ", num)
	}
}
