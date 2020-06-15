package main

import (
	"fmt"
)

func main() {
	var matrix = [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 9, 31, 32},
		{35, 36, 37, 38, 1002},
	}
	output := Search(matrix, 38)
	fmt.Println(output)
}

//Seach in sorted row and column matrix
func Search(matrix [][]int, target int) []int {
	row := 0
	col := len(matrix[0]) - 1
	for row <= len(matrix) && col >= 0 {
		if matrix[row][col] > target {
			col = col - 1
		} else if matrix[row][col] < target {
			row = row + 1
		} else {
			return []int{row, col}
		}

	}
	return []int{-1, -1}
}
