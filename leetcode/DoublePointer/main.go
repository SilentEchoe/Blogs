package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
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
