package main

import (
	"fmt"
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
