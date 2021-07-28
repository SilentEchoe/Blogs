package main

import (
	"fmt"
)

func main() {
	arry := []int{1, 3, 4,123,56, 7, 2, 1, 9}
	fmt.Println("冒泡排序：", bubbleSort(arry))
	fmt.Println("快速排序：",Quicksort(arry))
	fmt.Println("奇数偶数排序：",exchange(arry))
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

// 快速排序 
func Quicksort (arry []int)  []int {
	if len(arry) <= 1 {
		return arry
	}

	special := arry[0]
	mid := make([]int,0)
	hight := make([]int,0)
	low := make([]int,0)
	for i:=0; i< len(arry); i++ {
		if arry[i] == special {
			mid = append(mid,arry[i])
		}else if arry[i] < special {
			low = append(low,arry[i])
		}else if arry[i] > special {
			hight = append(hight,arry[i])
		}
	}
	low,hight = Quicksort(low), Quicksort(hight)
	myres := append(append(low,mid...),hight...)
	return myres
}

// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
/*
输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。
*/

func exchange(arry []int) []int {
	if len(arry) <= 1 {
		return arry
	}

	special := arry[0]
	mid := make([]int,0)
	hight := make([]int,0)
	low := make([]int,0)
	for i:=0; i< len(arry); i++ {
		if arry[i] == special {
			mid = append(mid,arry[i])
		}else if arry[i] %2 == 0 {
			hight = append(hight,arry[i])
		}else if arry[i] %2!=0 {
			low = append(low,arry[i])

		}
	}
	low,hight = exchange(low), exchange(hight)
	myres := append(append(low,mid...),hight...)
	return myres
}