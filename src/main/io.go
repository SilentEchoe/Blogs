package main

import (
	"bufio"
	"bytes"
	"fmt"
)



func main(){
	data := []byte("C语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)

	var buf [128]byte
	// [:] 代表全部
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)

	
}
