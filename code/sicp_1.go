package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var result string

	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()

	// 主线程继续执行逻辑
	fmt.Println("waiting for goroutine...")
	wg.Wait()

	fmt.Println("got result:", result)
}

func work() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("Job done")
		}
	}
}
