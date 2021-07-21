package main

import "fmt"

func main() {
	SliceTestOne()
}

// 这个面试提主要看两点
// 1. 切片的初始化是否合法
// 2. map中的顺序是随机的，不是有序的，这点需要注意
func SliceTestOne() {
	slice := []int{0, 1, 2, 3}

	m := make(map[int]*int)
	for k, v := range slice {
		m[k] = &v
	}

	for key, value := range m {
		fmt.Println(key, "->", *value)
	}

}
