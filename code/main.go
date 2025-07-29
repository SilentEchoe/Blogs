package main

import "fmt"

func main() {
	//fmt.Println(factorial(5))
	fmt.Println(countChange(100))
}

// SICP 练习1.11
// 函数f由如下规则定义：如果n<3，那么f(n)=n；如果n≥3，那么f(n)=f(n-1)+2f(n-2)+3f(n-3)。
// 请写一个函数，它通过一个递归计算过程计算f。再写一个函数，通过迭代计算过程计算f。
func f(n int) int {
	if n < 3 {
		return n
	}
	return f(n-1) + 2*f(n-2) + 3*f(n-3)
}

//func fib(n int) int {
//	if n == 0 {
//		return 0
//	}
//	if n == 1 {
//		return n
//	}
//	return fib(n-1) + fib(n-2)
//}

func fib(a, b, count int) int {
	if count == 0 {
		return b
	}
	return fib(a+b, a, count-1)
}

func fibIter(a) {

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

func countChange(amount int) int {
	return cc(amount, 5)
}

func cc(amount, kindsOfCoins int) int {
	if amount == 0 {
		return 1
	}
	if amount < 0 || kindsOfCoins == 0 {
		return 0
	}
	return cc(amount, kindsOfCoins-1) + cc(amount-firstDenomination(kindsOfCoins), kindsOfCoins)
}

func firstDenomination(kindsOfCoins int) int {
	switch kindsOfCoins {
	case 1:
		return 1
	case 2:
		return 5
	case 3:
		return 10
	case 4:
		return 25
	case 5:
		return 50
	default:
		return 0
	}
}
