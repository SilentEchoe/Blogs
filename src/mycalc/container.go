package main

import "fmt"

// Go 语言的容器用于各种形式的存储和处理数据的功能
// 类似于C# 中的 List

//声明数组 [4] 数组的长度为4
var list [4]string

//数组初始化
var q [3]int = [3]int{1, 2, 3}

// ...表示数组的长度是根据初始化值的个数来计算
q := [...]int{1, 2, 3}



func main() {
	list[0] = "第一个数"
	list[1] = "第二个数"
	list[2] = "第三个数"
	list[3] = "第四个数"

	for k, v := range list {
		fmt.Println(k, v)
	}

}
