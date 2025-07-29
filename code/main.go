package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

//func factorial(n int) int {
//	result := 1
//	for i := 2; i <= n; i++ {
//		result *= i
//	}
//	return result
//}
