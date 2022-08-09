package main

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
