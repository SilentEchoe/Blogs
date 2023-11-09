package main

func main() {
	//fmt.Println(selectionSort([]int{5, 3, 6, 2, 10}))
	selectSort([]int{5, 2, 3, 1})

}

// 查找数组中最小的元素
func findSmallest(sum []int) int {
	smallest := sum[0] //存储最小的值
	smallest_index := 0
	for i := 0; i < len(sum); i++ {
		if sum[i] < smallest {
			smallest = sum[i]
			smallest_index = i
		}
	}
	return smallest_index
}

// 选择排序

func selectionSort(sum []int) []int {
	newArr := make([]int, len(sum))

	var newsum = sum
	for i := 0; i < len(sum); i++ {
		// 拿到最小的那个元素的索引
		smallest := findSmallest(sum)
		// 放入新的数组中
		newArr = append(newArr, sum[smallest])
		// 旧数组中删除这个最小元素
		newsum = append(newsum[:smallest], newsum[smallest+1:]...)
	}
	return newArr
}

// 选择排序
func selectSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 已排序区间初始化为空，未排序区间初始化待排序切片
	for i := 0; i < len(nums); i++ {
		// 未排序区间最小值初始化为第一个元素
		min := i
		// 从未排序区间第二个元素开始遍历，直到找到最小值
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		// 将最小值与未排序区间第一个元素互换位置（等价于放到已排序区间最后一个位置）
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
}
