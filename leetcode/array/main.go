/*
	数组算法题

双指针技巧：主要分为两类，左右指针和快慢指针
左右指针就是两个相向或者相背而行。
快慢指针就是两个指针同向而行，一快一慢。

找回文串技巧：从中心向两端扩展的双指针技巧
如果回文串的长度为奇数，则它有一个中心字符；如果长度为偶数，则有两个中心字符串
*/
package main

import "fmt"

func main() {
	palindrome("123321", 0, 6)
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

// Code.27 移除元素
// 同Code.26 依然使用快慢指针的方法
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// Code.704 二分查找
// 标准使用数组二分查找
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		var mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return -1
}

// 反转数组
func reverse(sum []int) {
	left := 0
	right := len(sum) - 1

	for left < right {
		var tem = sum[left]
		sum[left] = sum[right]
		sum[left] = tem
		left++
		right--
	}
}

// 回文字符串

// palindrome 辅助函数
// 在s中寻找以 s[l] 和 s[r] 为中心的最长回文串
func palindrome(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		// 双指针，向两边展开
		left--
		right++
	}
	fmt.Println(s[left:right])
	return s[left:right]
}

// 最长回文串
func logestPalindrome(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)

		if len(res) < len(s1) {
			res = s1
		} else {
			res = res
		}
		if len(res) < len(s1) {
			res = s2
		} else {
			res = res
		}

	}
	return res
}
