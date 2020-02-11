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

// 为鸟添加Walk()方法,实现行走动物接口
func (b *bird) Walk()  {
	fmt.Println("bird:walk")
}

type pig struct {
}

// 为猪添加Walk()方法，实现行走动物接口
func (p *pig) Walk()  {
	fmt.Println("pig:walk")
}

//将接口转换为其他接口
func main()  {
	// 创建动物的名字到实例的映射
	animals := map[string]interface{}{
		"bird":new(bird),
		"pig" : new(pig),
	}


	for name,obj := range animals{
		f, isFlyer := obj.(Flyer)
		w, isWalker := obj.(Walker)

		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)

		if isFlyer  {
			f.Fly()
		}

		if isWalker {
			w.Walk()
		}

	}



}









