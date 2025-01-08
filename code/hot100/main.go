package main

/* LeetCode Hot 100 */

func main() {

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

// LeetCode
func groupAnagrams(strs []string) [][]string {
	return nil
}
