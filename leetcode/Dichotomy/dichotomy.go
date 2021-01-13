package main

func main() {
	s := []int{5, 7, 8, 9, 10, 11}
	var target = 9
	var a = search(s, target)
	var b = serachTwo(s, target)
	println(a)
	println(b)

	s1 := []int{1, 3}
	target2 := 3
	var c = searchInsert(s1, target2)
	println(c)
}

// 暴力查找
func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

// 二分法查找
func serachTwo(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start++
		} else {
			end--
		}
	}
	return -1
}

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//输入: [1,3,5,6], 5
//输出: 2

//输入: [1,3,5,6], 2
//输出: 1
func searchInsert(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}

		if nums[i] < target && target < nums[i+1] {
			return i + 1
		}

	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}

	return 0
}

// 二分法查询
func searchInsertTwo(nums []int, target int) int {

	return 0
}
