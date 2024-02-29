package main

func main() {
	println(binary_search([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4))
}

func binary_search(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
