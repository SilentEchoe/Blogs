package main

import "fmt"

func subsets(nums []int) []int {
	//result := make([][]int, 0)
	result := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Println("当前循环:", nums[i])
		for j := i; j >= 0; j-- {
			var d = j
			for d >= 0 {
				fmt.Print("种类有为：", nums[d])
				d--
			}

		}
		nums = nums[:len(nums)-1]
	}
	if len(nums) == 0 {
		return result
	}
	return nil
}

func main() {

	nums := make([]int, 0)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	subsets(nums)
}
