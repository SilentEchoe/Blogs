package main

// 动态规划算法的核心就是记住已经解决过的子问题的解
func main() {
	sums := []int{4, 2, 2, 3, 1, 4, 7, 8, 6, 9}
	magnet := make([]int, len(sums))
	subscript := 0
	for i := 0; i < len(sums)-1; i++ {
		temporary := sums[i]
		if i == 0 {
			for l := 0; l < i; l++ {
				if sums[l] > temporary {
					subscript = l
				}
			}
		}

		magnet = append(magnet, i)

	}

	println(subscript)

}
