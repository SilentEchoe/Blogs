package main

import "fmt"

func main() {
	var data = make([]int, 3)
	doWork(data)
	fmt.Println(data)
}

func doWork(data []int) {
	data = append(data, 1)
	data[0] = 1
}
