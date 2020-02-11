package main

import "io"

// 构造内嵌的例子

// 声明一个设备结构
type  device struct {

}

func (d *device) Write(p []byte) (n int,err error)  {
	return  0,nil
}

func (d *device)Close() error  {
	return  nil
}

func main()  {
	var wc io.WriteCloser = new(device)
	wc.Write(nil)
	wc.Close()
	var writeOnly io.Writer = new(device)
	writeOnly.Write(nil)
}



// 定义结构
type Writer interface {
	Write(p []byte) (n int, err error)
}


// 定义构造
type Closer interface {
	Closer() error
}

// 定义接口
type WriteCloser interface {
	Writer
	Closer
}
