package main

import "fmt"

func main() {
	var i int = 1
	var str string = "指针"

	fmt.Printf("%p %p", &i, &str)
}
