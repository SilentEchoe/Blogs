package main

import (
	"runtime"
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

	println(runtime.NumCPU())

	go NewDoTask()
	time.Sleep(10 * time.Microsecond)
	println("任务结束")

	select {}

}

func NewDoTask() {
	for i := 0; i < 10000; i++ {
		println("执行任务", i)
	}

}

func DoTask(m MyQueue) {

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
