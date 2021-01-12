package main

func main() {
	var sum = []int{5, 7, 7, 8, 8, 10}
	var sum2 = search_for_range(sum, 8)
	for v := range sum2 {
		println(v)
	}
}

//给定一个包含 n 个整数的排序数组，找出给定目标值 target 的起始和结束位置。
// 如果目标值不在数组中，则返回[-1, -1]
func search_for_range(sum []int, target int) []int {
	if len(sum) == 0 {
		return []int{-1, -1}
	}
	result := make([]int, 2)
	var start = 0
	var end = len(sum) - 1
	for start <= end {
		mid := (start + end) / 2
		if sum[mid] == target {
			result = append(result, mid)
		}
		if sum[mid] > target {
			end--
		} else {
			start++
		}
	}
	return result
}
