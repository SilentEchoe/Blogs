package main

import "fmt"

func main() {
	fmt.Println(doWork())
}

func doWork() bool {
	defer fmt.Println("输出一")
	defer fmt.Println("输出二")
	panic("输出三")
	return true
}
