// 类型断言的格式
// i代表接口变量,T代表转换的目标类型,t 代表转换后的变量。
// t := i.(T)

package main

import "fmt"


func main()  {
	source := []string{"A1","A2","A3","A4","A5"}

	slice := source[2:3:3]
	fmt.Println(slice)
	// 追加字段
	slice = append(slice,"kai")
	fmt.Println(slice)
}