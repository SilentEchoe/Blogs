package code

func main() {

}

// 练习1.12 通过递归生成一个帕斯卡三角形
func generatePascalTriangle(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}

	triangle := generatePascalTriangle(n - 1)
	prevRow := triangle[len(triangle)-1]
	newRow := make([]int, len(prevRow)+1)

	newRow[0] = 1
	new

}
