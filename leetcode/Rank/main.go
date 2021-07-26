package main

import (
	"fmt"
)

func main() {
	arry := []int{1, 3, 4, 7, 2, 1, 9}
	fmt.Println(bubbleSort(arry))
}

// 冒泡排序
func bubbleSort(arry []int) []int {
	for i := 0; i < len(arry)-1; i++ {
		for j := 0; j < len(arry)-1-i; j++ {

			if arry[j] > arry[j+1] {
				var temp = arry[j+1]
				arry[j+1] = arry[j]
				arry[j] = temp
			}

		}
	}
	return arry
}
