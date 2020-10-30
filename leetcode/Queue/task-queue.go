package main

type task struct {
	name  string
	param interface{}
}

func main() {
	// 队列
	var Queue []*task
	// 添加元素
	Queue = append(Queue, nil)
	// 移除首元素
	Queue = Queue[1:]

	var QueueLen = len(Queue)
	println(QueueLen)
}
