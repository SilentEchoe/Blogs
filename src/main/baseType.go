// 基本类型
package main


import (
    "fmt"
)

// 将int 定义为MyInt类型
type MyInt int

// 为基本类型添加方法
func (m MyInt) IsZero() bool  {
	return m == 0 
}

func main() {
	var b MyInt

	b = 1
    fmt.Println(b.IsZero())
}




