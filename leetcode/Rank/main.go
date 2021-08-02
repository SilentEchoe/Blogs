package main

import "fmt"

func main() {
	arry := []int{1, 3, 4, 123, 56, 7, 2, 1, 9}
	fmt.Println("冒泡排序：", bubbleSort(arry))
	fmt.Println("快速排序：", Quicksort(arry))

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

// 选择排序
func selectionSort(arry []int) []int {
	var len = len(arry)
	for i := 0; i < len-1; i++ {
		min := i

		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := 0; j < len-1; j++ {
			if arry[j] < arry[min] {
				min = j
			}
		}
		arry[i], arry[min] = arry[min], arry[i]

	}

	return arry
}

// 快速排序
func Quicksort(arry []int) []int {
	if len(arry) <= 1 {
		return arry
	}

	special := arry[0]
	mid := make([]int, 0)
	hight := make([]int, 0)
	low := make([]int, 0)
	for i := 0; i < len(arry); i++ {
		if arry[i] == special {
			mid = append(mid, arry[i])
		} else if arry[i] < special {
			low = append(low, arry[i])
		} else if arry[i] > special {
			hight = append(hight, arry[i])
		}
	}
	low, hight = Quicksort(low), Quicksort(hight)
	myres := append(append(low, mid...), hight...)
	return myres
}
