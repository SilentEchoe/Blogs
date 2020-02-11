package main

//i 代表接口变量，T 代表转换的目标类型，t 代表转换后的变量。
//t := i.(T)

import "fmt"

// 定义飞行动物接口
type  Flyer interface {
	Fly()
}

// 定义行走动物接口
type Walker interface {
	Walk()
}

// 定义鸟类
type bird struct {
}

// 实现飞行动物接口
func (b *bird) Fly()  {
	fmt.Println("bird:fly")
}

// 为鸟
