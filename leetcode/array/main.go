/*
	数组算法题

双指针技巧：主要分为两类，左右指针和快慢指针
左右指针就是两个相向或者相背而行。
快慢指针就是两个指针同向而行，一快一慢。
*/
package main

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	twoSum(numbers, target)
}

// Code.167 两数之和-输入有序数组
// nubers := []int{2, 7, 11, 15}
// target := 9
// 题解：使用双指针技巧，类似于二分查找法
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		var sum = numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			// 如果sum比target小，左指针++
			left++
		} else if sum > target {
			// 如果sum比target大，右指针向左偏移，所以--
			right--
		}
	}
	return []int{-1, -1}
}

// Code.26 删除有序数组中重复的项
// 输入：nums = [0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4]
// 题解：使用快慢指针
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
