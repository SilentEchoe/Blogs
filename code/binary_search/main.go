/* 二分查找 */
package main

func main() {

}

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target
//写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]
		if guess == target {
			return mid
		}
		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
