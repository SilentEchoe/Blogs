package main

import (
	"fmt"
	"time"
)

// ===================== 树形递归 =====================
var recursiveCallCount int

func fibRecursive(n int) int {
	recursiveCallCount++
	if n < 2 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// ===================== 线性迭代 =====================
func fibIterative(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// ===================== 记忆化递归 =====================
var memo = map[int]int{}
var memoCallCount int

func fibMemo(n int) int {
	memoCallCount++
	if n < 2 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibMemo(n-1) + fibMemo(n-2)
	return memo[n]
}

// ===================== 主程序 =====================
func main() {
	n := 40

	fmt.Println("==== 树形递归 ====")
	recursiveCallCount = 0
	start := time.Now()
	result := fibRecursive(n)
	elapsed := time.Since(start)
	fmt.Printf("fibRecursive(%d) = %d\n", n, result)
	fmt.Printf("调用次数: %d, 用时: %v\n\n", recursiveCallCount, elapsed)

	fmt.Println("==== 线性迭代 ====")
	start = time.Now()
	result = fibIterative(n)
	elapsed = time.Since(start)
	fmt.Printf("fibIterative(%d) = %d\n", n, result)
	fmt.Printf("用时: %v\n\n", elapsed)

	fmt.Println("==== 记忆化递归 ====")
	memoCallCount = 0
	memo = map[int]int{} // 清空缓存
	start = time.Now()
	result = fibMemo(n)
	elapsed = time.Since(start)
	fmt.Printf("fibMemo(%d) = %d\n", n, result)
	fmt.Printf("调用次数: %d, 用时: %v\n", memoCallCount, elapsed)

}

func sumIntegers(a, b int) int {
	if a > b {
		return 0
	}
	return a + sumIntegers(a+1, b)
}
