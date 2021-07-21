package main

import "fmt"

/*

 */

func main() {
	defer_call()
	fmt.Println(defer_callTwo())
	fmt.Println(defer_callThree())
}

func defer_call() {
	// 执行顺序是 先进后出 所以应该是 C B A
	defer func() { fmt.Println("A") }()
	defer func() { fmt.Println("B") }()
	defer func() { fmt.Println("C") }()
}

/*

 */
func defer_callTwo() int {
	var i int
	defer func() {
		i++
	}()

	defer func() {
		i++
	}()
	return i
}

func defer_callThree() (i int) {
	defer func() {
		i++
	}()

	return i
}
