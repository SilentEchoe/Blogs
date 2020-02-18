package main

import (
	"fmt"
	"runtime"
	"sync"
	)

func main()  {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// wg 用来等待程序完成
	// 计数加2,表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	//声明一个匿名函数 并创建goroutine

	go func() {
		// 在函数退出时候调用Done 来通知main 函数工作已经完成
		defer  wg.Done()

		for count := 0; count < 3 ; count++  {
			for char :='a'; char < 'a' +26 ; char++  {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		// 在函数退出时候调用Done 来通知main 函数工作已经完成
		defer  wg.Done()

		for count := 0; count < 3 ; count++  {
			for char :='A'; char < 'A' +26 ; char++  {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")



}