package main

import "fmt"

// 切片（slice） 为引用类型
// Go语言中切片的内部结构包含地址、大小和容量，切片一般用于快速地操作一块数据集合

//slice [开始位置 : 结束位置]

var highRiseBuilding [30]int

func main() {
	var a = [3]int{1, 2, 3}
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

}
