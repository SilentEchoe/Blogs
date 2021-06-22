package main

import "fmt"

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

func searchInsert(nums []int, target int) int {
	res := -1

	start := 0
	end := len(nums) - 1

	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}
	}

	if nums[start] >= target {
		return start
	}
	if nums[end] >= target {
		return end
	}

	if nums[end] < target {
		return end + 1
	}

	return res
}

func main() {
	var target = 2
	var nums = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, target))
}
