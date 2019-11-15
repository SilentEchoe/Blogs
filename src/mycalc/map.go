package main

import (
	"container/list"
	"fmt"
	"sync"
)

// Go语言中的Map 类似于C#的字典类型
var mapLit map[string]int

func main() {

	mapLit = map[string]int{"one": 1, "two": 2}
	fmt.Println(mapLit["one"])

	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	// 删除map里面的元素
	delete(scene, "brazil")

	for k, v := range scene {
		fmt.Println(k, v)
	}

	var scena sync.Map

	//添加元素
	scena.Store("greece", 97)
	scena.Store("london", 100)
	scena.Store("egypt", 200)

	// 从sync.map 中根据键取值
	fmt.Println(scena.Load("london"))

	// 删除键值对
	scena.Delete("london")
	// 遍历所有sync.Map中的键值对
	scena.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

	//sync.Map 没有提供获取 map 数量的方法，替代方法是在获取 sync.Map 时遍历自行计算数量，
	//sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。

	// list

	l := list.New()
	l.PushBack("fist")
	l.PushBack(67)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
