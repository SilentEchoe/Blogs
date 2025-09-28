package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor // 捕获 factor
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5)) // 10
	fmt.Println(triple(5)) // 15
}
