package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{12, 13, 4},
		{11, 14, 5},
		{10, 15, 6},
		{9, 8, 7},
	}
	fmt.Println(SpiralTraverse(matrix))
}

func SpiralTraverse(array [][]int) []int {
	rowStart := 0
	rowEnd := len(array) - 1
	colStart := 0
	colEnd :=  len(array[0]) -1
	arrayVal := make([]int,0)
	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart;i<= colEnd ; i++ {
			arrayVal = append(arrayVal,array[rowStart][i])
		}

		for i := rowStart + 1 ; i <= rowEnd ; i++ {
			arrayVal = append(arrayVal,array[i][colEnd])
		}

		for i := colEnd - 1;i>=colStart && rowStart<rowEnd  ;i--{
			arrayVal = append(arrayVal,array[rowEnd][i])
		}
		for i := rowEnd-1;i>= rowStart+1 && colStart < colEnd ;i-- {
			arrayVal = append(arrayVal,array[i][colStart])
		}

		rowStart++
		rowEnd--
		colStart++
		colEnd--
	}
	return arrayVal
}

