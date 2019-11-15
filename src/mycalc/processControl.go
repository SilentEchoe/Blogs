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

	//Go 语言中的循环语句只支持for关键字
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i

	}
	fmt.Println(sum)

	for {
		sum++
		if sum > 100 {
			break
		}
	}

	fmt.Println(sum)

	// Go语言的 for 循环同样支持 continue 和 break 来控制循环
	// 但是它提供了一个更高级的 break，可以选择中断哪一个循环

JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Println(i)
		}
	}

	b := "hello"
	switch b {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	switch b {
	// 一分支多值
	case "mum", "hello":
		fmt.Println(1)
	}

	//跨越 case 的 fallthrough——兼容C语言的 case 设计
	// 执行完一个独立的case 分支后不会结束，会继续向下执行
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}

}
