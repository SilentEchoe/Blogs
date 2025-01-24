package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// 创建一个 done channel，用于发送关闭信号
	done := make(chan struct{})

	// 监听系统信号 Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	rule := make(chan int, 10)

	go producer(rule, done)
	go work(rule, done)

	fmt.Println("worker started, press ctrl+c to terminate...")
	<-sigChan

	fmt.Println("EdgeExporter worker closing...")
	// 发送关闭信号
	close(done)

	// 等待一段时间，确保生产者和消费者有足够的时间退出
	time.Sleep(time.Second * 5)

	fmt.Println("EdgeExporter worker safe exit...")
}

func producer(rule chan int, done chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			for i := 0; i < 10; i++ {
				rule <- i
			}
		}
	}
}

func work(rule chan int, done chan struct{}) {
	var wg sync.WaitGroup
	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(rule chan int, done chan struct{}) {
			for {
				select {
				case <-done:
					wg.Done()
					return
				case val, ok := <-rule: // 从 channel 接收数据:
					if !ok { // channel 已关闭
						fmt.Printf("消费者检测到 channel 关闭，退出")
						return
					}
					fmt.Printf("消费者 %d 消费了: %d\n", val)
					time.Sleep(time.Millisecond * 500) // 模拟消费耗时
				}
			}

		}(rule, done)
	}
}
