// 类型断言的格式
// i代表接口变量,T代表转换的目标类型,t 代表转换后的变量。
// t := i.(T)

package main

import "fmt"


func main()  {
	var colors map[string]string

	colors["Red"]="Red"

	source := []string{"A1","A2","A3","A4","A5"}

	// 截取数组的第三个元素
	slice := source[2:3:3]
	fmt.Println(slice)
	// 追加字段
	slice = append(slice,"kai")
	for index,value := range slice {
		fmt.Printf("Index: %d value: %d\n",index,value)

	}
}