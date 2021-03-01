package main

// 动态规划算法的核心就是记住已经解决过的子问题的解
// 查找数组中：左边的数全都小于它，右边数全都大于它
func main() {
	sums := []int{4, 2, 2, 3, 1, 4, 7, 8, 6, 9}
	poleSums := magnet_pole_sum(sums)

	for p := range poleSums {
		println(p)
	}
}

func magnet_pole_sum(sum []int) []int {
	magnet := make([]int, 0)
	leftMaxSum := 0
	for i := 0; i < len(sum); i++ {
		//println("当前参数为：", sum[i])
		if sum[i] > leftMaxSum {
			leftMaxSum = sum[i]
		}

		for r := i + 1; r < len(sum)-2; r++ {
			println("当前参为:", sum[i])
			println("右参为：", sum[r])

			if sum[r] > sum[i] && sum[i] > leftMaxSum {
				println(sum[i])
				//magnet = append(magnet, sum[i])
			}
		}

	}

	return magnet
}
