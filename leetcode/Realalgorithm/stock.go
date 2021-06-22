package Realalgorithm

import "fmt"

// 计算股票跨度
// 输入 保存n个股票报价的数组
// 输出 保存n个股票跨度的数组

func SimpleStockSpan(quotes []int) []int {
	spanl := make([]int, len(quotes))
	for i := 0; i <= len(quotes)-1; i++ {
		var k = 1
		var span_end = false

		for i-k >= 0 && !span_end {
			if quotes[i-k] <= quotes[i] {
				k += 1
			} else {
				span_end = true
			}
		}
		spanl[i] = k
	}
	return spanl
}

func main() {
	quote := []int{7, 11, 8, 7, 3, 9, 11}
	spanl := SimpleStockSpan(quote)
	for _, v := range spanl {
		fmt.Println(v)
	}
}
