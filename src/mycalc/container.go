package main

import (
	"fmt"
)

// Go 语言的容器用于各种形式的存储和处理数据的功能
// 类似于C# 中的 List

//声明数组 [4] 数组的长度为4
var list [4]string

//数组初始化
//var q [3]int = [3]int{1, 2, 3}

// ...表示数组的长度是根据初始化值的个数来计算
//q := [...]int{1, 2, 3}

// 多维数组
// 声明一个多维数组，两个维度的长度分别是4和 2
var array [4][2]int

// 初始化二维数组

var arrayList [2][2]int

func main() {
	list[0] = "第一个数"
	list[1] = "第二个数"
	list[2] = "第三个数"
	list[3] = "第四个数"

	array = [4][2]int{{10, 11}, {20, 21}, {30}, {40}}

	for k, v := range list {
		fmt.Println(k, v)
	}
	arrayList[0][0] = 10
	fmt.Println(array[0][1])
	fmt.Println(arrayList[0][0])

}
