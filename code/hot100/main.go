package main

import (
	"sort"
)

/* LeetCode Hot 100 */

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

// LeetCode.1 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, v}
		}
		m[v] = i
	}
	return nil
}

// LeetCode.49 字母异位词分组
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}
	ans := make([][]string, 0, len(m))

	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}

// LeetCode.128 最长连续序列
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	// 先把数组里面所有的元素放入map
	for _, num := range nums {
		m[num] = true
	}
	sum := 0
	for num := range m {
		if !m[num-1] {
			currentNum := num
			currentStreak := 1
			// 扩展连续序列
			for m[currentNum+1] {
				currentNum++
				currentStreak++
			}
			// 更新最长连续序列的长度
			if currentStreak > sum {
				sum = currentStreak
			}
		}
	}
	return sum
}

// LeetCode.283 移动零
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func moveZeroes(nums []int) {
	//使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。
	//右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。
	//注意到以下性质：
	//左指针左边均为非零数；
	//右指针左边直到左指针处均为零。
	//因此每次交换，都是将左指针的零与右指针的非零数交换，且非零数的相对顺序并未改变。
	left, right, n := 0, 0, len(nums)
	for right < n {
		//当右指针不为零
		if nums[right] != 0 {
			// 交换
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
