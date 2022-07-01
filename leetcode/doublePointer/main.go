package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
	rotateNums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(rotateNums, 3)
	fmt.Println(rotateNums)

	str := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(str)

	reverseWords("Let's take LeetCode contest")
}

//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = WitBranch(nums[i])
		nums[i] = nums[i] * nums[i]

	}
	sort.Ints(nums)
	return nums

}

func WitBranch(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

/*-----------------双指针-----------------------------------------*/

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。
//思路：移除所有0 直接后面补
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

//编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
//不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
//你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
//输入：["h","e","l","l","o"]
//输出：["o","l","l","e","h"]

func reverseString(s []byte) string {

	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
	return string(s)
}

func reverseWords(s string) string {
	string_slice := strings.Split(s, " ")
	var buffer bytes.Buffer
	for _, v := range string_slice {
		var data []byte = []byte(v)
		newString := reverseString(data)
		buffer.WriteString(newString)
		buffer.WriteString(" ")
	}
	s = buffer.String()
	s = strings.TrimRight(s, ",")

	return s
}
