package main

import "fmt"

// 切片（slice） 为引用类型
// Go语言中切片的内部结构包含地址、大小和容量，切片一般用于快速地操作一块数据集合

//slice [开始位置 : 结束位置]

var highRiseBuilding [30]int

// 切片类型可直接声明 语法如下：
// 声明字符串切片
//var strList []string

func main() {
	var a = []int{1, 2, 3}
	// a[0:1] 代表从第一个元素开始 到第二个元素结束
	fmt.Println(a, a[0:1])

	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}

	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])

	// 使用Make()函数构造切片
	// 语法为 make( []Type, size, cap ) size指的是这个类型分配多少个元素，Cap 为预分配的元素数量
	// 这个值不影响size，只能提前分配空间，降低多次分配空间造成的性能问题
	c := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(c, b)
	fmt.Println(len(c), len(b))

	//为数组a 添加追加一个元素
	a = append(a, 1)
	fmt.Println(a)

	// 追加多个元素, 手写解包方式
	a = append(a, 1, 2, 3)
	// 追加一个切片
	a = append(a, []int{1, 2, 3}...)

	//在开头添加一个元素
	a = append([]int{4}, a...)

	// 在开头添加1个切片
	a = append([]int{-3, -2, -1}, a...)

	//在切片开头添加元素一般都会导致内存重新分配，而且会导致已有元素全部被复制一次
	// 因此在切片的开头添加元素的性能要比从尾部追加元素的性能要差很多

	// 删除开头第一个元素
	a = a[1:]

	// 删除开头1个元素 使用append进行删除
	a = append(a[:0], a[1:]...)

	// 删除中间的元素

	a = append(a[:2], a[2+1:]...)

	// 删除尾部删除
	a = a[:len(a)-1]

	//range 关键词

	for index, value := range a {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	fmt.Println(a)
}
