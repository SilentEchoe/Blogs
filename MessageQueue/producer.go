package main

import "fmt"

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

var done = make(chan int)

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

	//println(runtime.NumCPU())

	//go NewDoTask()
	go DoTask(globalQueue)

	cancelled(done)
	//time.Sleep(20 * time.Second)
	println("任务结束")

}

func cancelled(done chan int) {
	select {
	case <-done:
		fmt.Println("quit")
		return
	}
}

func NewDoTask() {
	for i := 0; i < 100; i++ {
		println("执行任务", i)
	}
	done <- 0
}

func DoTask(m []MyQueue) {
	for _, value := range m {
		if value.key == "show" {
			//time.Sleep(1000)
			println("执行show 任务")
		}

		if value.key == "update" {
			//time.Sleep(1000)
			println("update 任务")
		}
	}
	done <- 0

}

func showQueQu() {

	for _, value := range globalQueue {
		println(value.key, value.value)
	}
}
