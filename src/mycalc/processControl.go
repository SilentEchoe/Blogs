package main

import "fmt"

func main() {
	a := 1

	// 判断a的值是否等于2 如果不等于2
	if !(a == 2) {
		fmt.Println("a不等于2")
	} else {
		fmt.Println("a 等于2")
	}
}
