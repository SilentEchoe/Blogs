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

	// if的特殊写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断

	//if err := Connect(); err != nil {
	//	fmt.Println(err)
	//	return
	//}

}
