package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 9, 12}
	fmt.Println(search(nums, 9))
	fmt.Println(searchFirstBadVersion(5))
	numTwos := []int{1, 2}
	fmt.Println("寻找峰值", findPeakElement(numTwos))

	numbers := []int{0, 0, 3, 4}
	target := 0
	fmt.Println("寻找合并值:", twoSum(numbers, target))

	searchRangenums := []int{5, 7, 7, 8, 8, 10}
	searchRange := 8
	fmt.Println(searchRangeTwo(searchRangenums, searchRange))
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
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[mid+1] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

//输入：numbers = [2,7,11,15], target = 9
//输出：[1,2]
//解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

func twoSum(numbers []int, target int) []int {
	if len(numbers) <= 0 {
		return nil
	}

	start := 0
	end := len(numbers) - 1

	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			res := make([]int, 0)
			res = append(res, start+1)
			res = append(res, end+1)
			return res
		} else if sum < target {
			start++
		} else {
			end--
		}
	}
	return nil
}

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
func searchRangeTwo(nums []int, target int) []int {
	start := 0
	end := len(nums) - 1
	res := make([]int, 0)
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			res = append(res, mid)
			start = mid + 1
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}
	}

	if len(res) == 0 {
		res = append(res, -1)
		res = append(res, -1)
	}
	return res
}
