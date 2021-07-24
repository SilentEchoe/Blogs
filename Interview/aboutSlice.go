package main

import "fmt"

func main() {
	SliceDilatation()
	SliceTestOne()
}

// SliceTestOne 这个面试提主要看两点
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

// SliceAdding 切片和切片相加
func SliceAdding() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

// SliceDilatation 切片扩容
// 这要主要是的 append 方法是在函数内部,底层数组和外面的不是一个，所以没有改变外面的。
func SliceDilatation() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)
}

func ap(a []int) {
	a = append(a, 10)
}

func app(a []int) {
	a[0] = 1
}
