/*
二分查找算法
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	count := 0
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
