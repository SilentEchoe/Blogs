package main

import "fmt"

func main() {
	SliceTestOne()
}

// 这个面试提主要看两点
// 1. for range 循环的时候会创建每个元素的副本,不是每个元素的引用，所以取值时都是变量最后赋的值
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

// 切片和切片相加
func SliceAdding() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
