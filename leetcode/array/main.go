package main

//只要数组有序，就应该想到双指针技巧

func main() {

}

// 快慢指针技巧
// 删除有序数组中的重复项
// 快指针走前面，找到一个不重复的元素就赋值给 slow 让 slow 前进一步（因为是有序数组
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	fast := 0

	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			// 慢指针走一步
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// 移除元素
func removeElement(nums []int, val int) []int {
	if len(nums) == 0 {
		return nil
	}

	slow := 0
	fast := 0

	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return nums
}

// 左右指针(二分查找法示例
func binarySearch(nums []int, target int) int {
	// 一左一右两个指针相向而行
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 中线
		mid := (right + left) / 2
		if nums[mid] == target {
			// 如果查到了直接返回
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

// 两数之和
// 使用二分查找来解决
func twoSum(nums []int, target int) (int, int) {
	left := 0
	right := len(nums) - 1
	for left <= right {
		sum := nums[left] + nums[right]
		if sum == target {
			return left + 1, right + 1
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return -1, -1
}

// 最长回文子串
// 回文子串问题则是让左右指针从中心向两端扩展
