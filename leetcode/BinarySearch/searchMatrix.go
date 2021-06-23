package main

import (
	"fmt"
)

//编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-a-2d-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func searchMatrix(matrix [][]int, target int) bool {
	var row = len(matrix)
	var column = len(matrix[0])

	newMatrix := make([]int, 0)
	for i := 0; i < row; i++ {
		for c := 0; c < column; c++ {
			newMatrix = append(newMatrix, matrix[i][c])
		}
	}

	start := 0
	end := len(newMatrix) - 1

	for start+1 < end {
		mid := start + (end-start)/2
		if newMatrix[mid] == target {
			return true
		} else if newMatrix[mid] < target {
			start = mid
		} else if newMatrix[mid] > target {
			end = mid
		}
	}

	if newMatrix[start] == target {
		return true
	}

	if newMatrix[end] == target {
		return true
	}

	return false
}

func main() {

	matrix := [][]int{{1}}

	var target = 0
	fmt.Println(searchMatrix(matrix, target))
}
