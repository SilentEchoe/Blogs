package main

import "fmt"

func main() {
	fmt.Println(quicksort([]int{10, 5, 2, 3}))
}

// 快速排序
// 从数列中挑出一个元素，作为基准（pivot）
// 重新排序数列，所有比基准小的值放到基准前面，所有比基准大的值放到基准后面。排序之后，基准值便处于数列的中间位置，这个过程称为分区。
func quicksort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	pivot := arr[0] // 选出基准
	var less []int
	for i := 1; i < len(arr); i++ {
		if i < pivot {
			less = append(less, arr[i])
		}
	}

	var greater []int
	for i := 1; i < len(arr); i++ {
		if i > pivot {
			greater = append(greater, arr[i])
		}
	}

	var newarr = append(quicksort(less))
	newarr = append(newarr, pivot)
	newarr = append(quicksort(greater))

	return newarr
}
