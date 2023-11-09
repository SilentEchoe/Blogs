package main

func main() {
	//fmt.Println(selectionSort([]int{5, 3, 6, 2, 10}))
	selectSort([]int{5, 3, 6, 2, 10})
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
func selectSort(sum []int) {
	if len(sum) == 0 {
		return
	}

	for i := 0; i < len(sum)-1; i++ {
		min := i

		for j := i + 1; j < len(sum); j++ {
			if sum[j] < sum[min] {
				min = j
			}
			if min != j {
				// 交换元素
				sum[i], sum[min] = sum[min], sum[i]
			}
		}

	}
}
