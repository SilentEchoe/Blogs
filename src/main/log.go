// 类型断言 用于检查接口类型变量所持有的值是否实现期望的接口 或具体的类型

//value,ok := x.(T)
package main

import "fmt"

func  main()  {
	var a int
	a = 10
	// 验证类型
	getType(a)
	
}


func getType(a interface{}){
	switch a.(type){
		case int:
			fmt.Println("the type of a is int")
		case string:
			fmt.Println("the type of a is string")
		case float64:
			fmt.Println("the type of a is float64")
		default:
			fmt.Println("unknown type")
	}
}