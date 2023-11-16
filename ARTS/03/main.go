package main

import "fmt"

func main() {
	list := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	quicksort(list, 0, len(list)-1)
	fmt.Println(list)
}

// 快速排序
// 从数列中挑出一个元素，作为基准（pivot）
// 重新排序数列，所有比基准小的值放到基准前面，所有比基准大的值放到基准后面。排序之后，基准值便处于数列的中间位置，这个过程称为分区。
func quicksort(list []int, low, high int) {
	if high > low {
		//位置划分
		pivot := partition(list, low, high)
		//左边部分排序
		quicksort(list, low, pivot-1)
		//右边排序
		quicksort(list, pivot+1, high)
	}
}

// 基准值
func partition(list []int, low, high int) int {
	pivot := list[low]
	for low < high {
		for low < high && pivot <= list[high] {
			high--
		}
		list[low] = list[high]

		for low < high && pivot >= list[low] {
			low++
		}
		list[high] = list[low]
	}
	list[low] = pivot
	return low
}
