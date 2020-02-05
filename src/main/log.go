// 类型断言 用于检查接口类型变量所持有的值是否实现期望的接口 或具体的类型

//value,ok := x.(T)


func  main()  {
	// 申明一个接口
	var x interface{}
	x = 10
	value, ok := x.(int)

	fmt.Print(value, ",", ok)

}