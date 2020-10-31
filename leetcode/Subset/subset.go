package main

func subsets(nums []int) [][]int {
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)

	if len(nums) == 0 {
		println("数组为空")
		return nil
	}

	for i := len(nums) - 1; i >= 0; i-- {
		println(nums[i])
	}

	return nil
}

func main() {
	subsets(nil)
}
