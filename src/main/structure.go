package main

import (
	"fmt"
)

// Address构造函数
type Address struct {
    Province    string
    City        string
    ZipCode     int
    PhoneNumber string
}

// Bag 结构体
type Bag struct {
    items []int
}

// Insert 为定义新增方法
//   b *Bag 为接收器 
func (b *Bag) Insert(itemid int) {
    b.items = append(b.items, itemid)
}

func main()  {
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}
	fmt.Println(addr)
	b := new(Bag)
    b.Insert(1001)
	fmt.Println(b)
}

