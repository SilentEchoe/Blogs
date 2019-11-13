package main

//package main

import (
	"fmt"
	"strconv"
)

type Weapon int

// 将NewInt定义为int 类型
type NewInt int

// 将int取一个别名叫 IntAlias
type IntAlias = int

const (
	Arrow Weapon = iota // 开始生成枚举
	Shuriken
	SniperRifle
	Rifle
	Blower
)

var a string = "3"

func main() {
	// strconv.Itoa() 将int 类型转换为string 类型
	fmt.Println("a" + strconv.Itoa(32)) // a32
	fmt.Println(strconv.Atoi(a))
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)

	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon)

}
