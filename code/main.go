package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(100 * time.Second)
			fmt.Println("Main Task")
		}
	}()

	go work()

	// 使用通道阻塞，防止程序立即退出
	ch := make(chan struct{})
	<-ch
}

func work() {
	var wg sync.WaitGroup
	var rw sync.RWMutex
	var rule []string

	// 初始化一次随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生产者协程
	go func() {
		for {
			// 五秒写一次
			rw.Lock()
			newRule := writeRule()
			rule = newRule
			rw.Unlock()
			time.Sleep(5 * time.Second)
		}
	}()

	// 消费者
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				rw.RLock()
				fmt.Println(rule)
				rw.RUnlock()
				time.Sleep(1 * time.Second)
			}
		}()
	}
	// 等待一段时间，观察输出，实际应用中可以使用更优雅的方式如信号或退出通道
	time.Sleep(30 * time.Second)
}

func writeRule() []string {
	sliceLength := 10
	strSlice := make([]string, sliceLength)
	for i := 0; i < sliceLength; i++ {
		// 生成一个长度为 5 的随机字符串
		str := ""
		for j := 0; j < 5; j++ {
			// 随机生成字符的 ASCII 码，范围为 97 到 122 (a 到 z)
			char := byte(rand.Intn(26) + 97)
			str += string(char)
		}
		strSlice[i] = str
	}
	return strSlice
}
