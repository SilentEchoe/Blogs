package main

import (
	"fmt"
)

// 变量、指针和地址三者的关系是，每个变量都拥有地址，指针的值就是地址。
func main() {
	var cat int =1 
	var str string = "banana"
	fmt.Printf("%p %p",&cat,&str)


}
