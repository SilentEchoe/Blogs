package main

import (
	"fmt"
)

func main() {

	//ch := make(chan int, 10)
	//go producer(ch)
	//for c := range ch {
	//	fmt.Println(c)
	//}
	//time.Sleep(2000)

}

func producer(c chan int) {
	defer fmt.Println("生产任务结束")
	for i := 1; i <= 10; i++ {
		fmt.Println("生产任务:", i)
		c <- i
	}
	close(c)
}
