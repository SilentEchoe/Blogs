package main

import "fmt"

func main() {
	mate := []string{"1", "2", "3"}
	mateInfo(mate)
}

func mateInfo(args ...any) {
	fmt.Println(args)
}
