// 变量逃逸分析
package main


import (
    "fmt"
)

func main() {
	// 变量a 逃逸到堆
	var a int
	void()
	// dummy(0)逃逸到堆
	fmt.Println(a,dummy(0)) 
}


func dummy(b int) int{
	// 声明一个变量C并赋值
	var c int
	c = b
	return c
}

func void(){

}

