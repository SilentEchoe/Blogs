package main

import "fmt"

func main() {
	compareStruct()
}

// 结构体比较
func compareStruct() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "11"}
	if sn1 == sn2 {
		fmt.Println("能比较")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	_ = sm1
	_ = sm2
	// 结构体中包含 Slice map 函数 则不能比较
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}
}
