package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 9, 12}
	fmt.Println(search(nums, 9))
	fmt.Println(searchFirstBadVersion(5))
	numTwos := []int{1, 2}
	fmt.Println("寻找峰值", findPeakElement(numTwos))
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}

	}
	if nums[start] == target {
		return start
	}

	if nums[end] == target {
		return end
	}
	return -1
}

func firstBadVersion(n int) int {
	nums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		nums = append(nums, 1)
	}
	return search(nums, 4)
}

func searchFirstBadVersion(n int) int {
	start := 1
	end := n
	for start < end {
		mid := start + (end-start)/2

		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}

//输入：nums = [1,2,3,1]
//输出：2
//解释：3 是峰值元素，你的函数应该返回其索引 2。

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
