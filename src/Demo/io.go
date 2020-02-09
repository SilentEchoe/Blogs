package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main(){
	data := []byte("C语言中文网")
	// 读取并返回一个字节，如果没有字节可读，则返回错误信息
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)


	var buf [128]byte
	// [:] 代表全部
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	Get()
}

func Get(){
	fmt.Println("输入")
}
