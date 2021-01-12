package main

func main() {
	s := []int{5}
	var target = 5
	var a = search(s, target)
	println(a)
}

func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
