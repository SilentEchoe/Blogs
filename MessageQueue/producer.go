package main

type MyQueue struct {
	message string
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
	var task1 = MyQueue{
		message: "task1",
		value:   "show version",
		IsExist: 0,
	}
	var task2 = MyQueue{
		message: "task2",
		value:   "show log",
		IsExist: 0,
	}

	QueQuSend(task1)
	QueQuSend(task2)
	showQueQu()
	QueQuPopup()
	showQueQu()

}

func showQueQu() {
	for _, value := range globalQueue {
		println(value.message, value.value)
	}
}
