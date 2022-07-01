package main

import "fmt"

//描述
//给定一个包含 n 个整数的排序数组，找出给定目标值 target 的起始和结束位置。
//
//如果目标值不在数组中，则返回[-1, -1]

/**
 * @param A: an integer sorted array
 * @param target: an integer to be inserted
 * @return: a list of length 2, [index1, index2]
 */
func searchRange(A []int, target int) []int {
	res := make([]int, 0)
	start := 0
	end := len(A) - 1

	if A[start] == target {
		res = append(res, start)
	}

	if A[end] == target {
		res = append(res, end)
	}

	if len(res) == 2 {
		return res
	}

	for start+1 < end {
		mid := start + (end-start)/2
		if A[mid] < target {
			start = mid
		} else if A[mid] == target {
			res = append(res, mid)
			start = mid
		} else if A[mid] > target {
			end = mid
		}

	}

	if len(res) == 1 {
		res = append(res, res[0])
	}

	if len(res) < 2 {
		res = append(res, -1)
		res = append(res, -1)
	}

	return res
}

func main() {
	var target = 2
	var A = []int{-1, 0, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 6, 90, 92, 93, 101}
	res := searchRange(A, target)
	for _, v := range res {
		fmt.Println(v)
	}
}
