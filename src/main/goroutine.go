package main

import (
	"fmt"
	"runtime"
	"sync"
	)

var wg sync.WaitGroup

func main()  {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU())

	// wg 用来等待程序完成
	// 计数加2,表示要等待两个goroutine

	wg.Add(2)

	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Fininsh")
	wg.Wait()

	fmt.Println("Terminating Program")

}

func printPrime(prefix string)  {
	// 在函数退出时调用Done来通知main 函数已经
	defer  wg.Done()

netx:
	for outer := 2; outer<5000 ;outer++  {
		for inner := 2;inner < outer ;inner++  {
			if outer%inner == 0 {
				continue netx
			}
		}
		fmt.Printf("%s:%d\n",prefix,outer)
	}
	fmt.Println("Compeleted",prefix)
}
