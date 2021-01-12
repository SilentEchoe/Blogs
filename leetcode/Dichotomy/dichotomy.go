package main

func main() {
	s := []int{5, 7, 8, 9, 10, 11}
	var target = 9
	var a = search(s, target)
	var b = serachTwo(s, target)
	println(a)
	println(b)
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
