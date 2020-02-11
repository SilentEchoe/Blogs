package main

import (
	"fmt"
)

// 定义一个接口
//type class interface{
//	Write([]byte) error
//}

type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构,用于实现DataWriter
type file struct {
}

type Log struct {
}

// 实现DataWriter接口的 WriteData
func (d *file)  WriteData(data interface{}) error {
	fmt.Println("WriteData",data)
	return nil
}

func (d *Log)  WriteData(data interface{}) error {
	fmt.Println("WriteData",data)
	return nil
}


func main()  {
	// 实例一个构造
	f := new(file)
	//声明一个接口
	var writer DataWriter
	writer = f
	writer.WriteData("data")

	l := new(Log)
	//声明一个接口
	var writer2 DataWriter
	writer2 = l
	writer2.WriteData("data")

}

