package main

import "fmt"

// Go语言中的Map 类似于C#的字典类型
var mapLit map[string]int

func main() {

	mapLit = map[string]int{"one": 1, "two": 2}
	fmt.Println(mapLit["one"])

	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for k, v := range scene {
		fmt.Println(k, v)
	}

}
