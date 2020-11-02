package main

import "fmt"

func subsets(nums []int) []int {

	//result := make([][]int, 0)
	result := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Println("参数为:", nums[i])
		for j := i - 1; j >= 0; j-- {
			fmt.Println("参数为", nums[i], nums[j])
		}
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
