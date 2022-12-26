/*
	数组算法题

双指针技巧：主要分为两类，左右指针和快慢指针
左右指针就是两个相向或者相背而行。
快慢指针就是两个指针同向而行，一快一慢。

找回文串技巧：从中心向两端扩展的双指针技巧
如果回文串的长度为奇数，则它有一个中心字符；如果长度为偶数，则有两个中心字符串

数组前缀和：
使用一个新的数组 preSum，然后用preSum[i] 记录 num[0...i-1]的累计加

差分数组：
和前缀和思想非常类似，差分数组的主要适用场景是频繁对原始数组对某个区间对元素进行增减
*/
package main

import "fmt"

func main() {
	NumArray([]int{3, 5, 2, -2, 4, 1})
	sum := sumRange(0, 3)
	fmt.Println(sum)
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

// 数组前缀和：
// 使用一个新的数组 preSum，然后用preSum[i] 记录 num[0...i-1]的累计加
var preSum = make([]int, 10)

func NumArray(nums []int) {
	for i := 1; i < len(nums)+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
}

func sumRange(left int, right int) int {
	return preSum[right+1-preSum[left]]
}

// 二维矩阵中的前缀和
// 定义 preSum[i][j] 记录 matrix 中子矩阵 [0,0,i-1,j-1] 的元素和
var preSumMatrix = make([][]int, 10)

func NumMatrix(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	if m == 0 || n == 0 {
		return
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			preSumMatrix[i][j] = preSumMatrix[i-1][j] + preSumMatrix[i][j-1] + matrix[i-1][j-1] - preSumMatrix[i-1][j-1]
		}
	}

}

// 计算子矩阵 [x1,y1.x2,y2]
func sumRegion(x1 int, y1 int, x2 int, y2 int) int {
	return preSumMatrix[x2+1][y2+1] - preSumMatrix[x1][y2+1] - preSumMatrix[x2+1][y1] + preSumMatrix[x1][y1]
}

// [差分数组]技巧
// 差分数组的主要使用场景是频繁对原始数组对某个区间的元素进行增减

var diff []int

// 构造差分数组

func Difference(nums []int) {
	if len(nums) == 0 {
		return
	}
	diff[0] = nums[0]
	for i := 0; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
}

func increment(i int, j int, val int) {
	diff[i] += val
	if j+1 < len(diff) {
		diff[j+1] -= val
	}
}

func result() []int {
	var res []int
	res[0] = diff[0]
	for i := 0; i < len(diff); i++ {
		res[i] = res[i-1] + diff[i]
	}
	return res
}
