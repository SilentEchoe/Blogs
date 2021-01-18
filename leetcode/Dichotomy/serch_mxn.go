package main

import (
	"fmt"
)

func main() {
	var matrix = [][]int{
		{1}, {3},
	}
	var target = 3
	var isexist = searchMatrix(matrix, target)
	fmt.Println(isexist)
}

func searchMatrix(matrix [][]int, target int) bool {
	var row = len(matrix)
	var column = len(matrix[0])
	for i := 0; i < row; i++ {
		if matrix[i][column-1] == target {
			return true
		}

		if target < matrix[i][column-1] {
			for j := 0; j < column; j++ {
				if matrix[i][j] == target {
					return true
				}
			}
			return false
		}
	}

	return false
}
