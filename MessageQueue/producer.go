package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type MyQueue struct {
	key     string
	value   string
	IsExist int
}

var globalQueue = make([]MyQueue, 0)

func QueQuSend(queue MyQueue) {
	globalQueue = append(globalQueue, queue)
}

func QueQuPopup() {
	globalQueue = globalQueue[1:]
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		var task1 = MyQueue{
			key:     "show",
			value:   "show log",
			IsExist: 0,
		}
		QueQuSend(task1)
		var task2 = MyQueue{
			key:     "update",
			value:   "show log",
			IsExist: 0,
		}
		QueQuSend(task2)
	}

	runtime.GOMAXPROCS(1)
	wg.Add(1)

	go NewDoTask()
	wg.Wait()
	//go NewDoTask()

	println("任务结束")
}

func NewDoTask() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		println("执行任务", i)
	}

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func DoTask(m MyQueue) {
	defer wg.Done()

	if m.key == "show" {
		//time.Sleep(1000)
		println("执行show 任务")
	}

	if m.key == "update" {
		//time.Sleep(1000)
		println("update 任务")
	}
}

func showQueQu() {

	for _, value := range globalQueue {
		println(value.key, value.value)
	}
}
