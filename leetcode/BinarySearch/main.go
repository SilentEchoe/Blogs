package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 9, 12}
	fmt.Println(search(nums, 9))

	fmt.Println(searchFirstBadVersion(5))
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
