package main

import (
	"fmt"
	"strings"
)

func main() {
	var rec = "12=13"
	k, v, ok := strings.Cut(rec, "=")
	if !ok {
		fmt.Println("解析失败")

	}
	fmt.Println(k)
	fmt.Println(v)
}
