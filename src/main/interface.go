package main

import "fmt"

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

// 实现DataWriter接口的 WriteData
func (d *file)  WriteData(data interface{}) error {
	fmt.Println("WriteData",data)
	return nil
}

func main()  {
	f := new(file)

	//声明一个接口
	var writer DataWriter

	writer = f

	writer.WriteData("data")


}

