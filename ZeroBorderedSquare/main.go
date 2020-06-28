package main

import "fmt"

func main() {
	matr := [][]int{
		{0, 0, 0},
		{1, 0, 0},
		{0, 0, 0},
	}
	for i:=0;i<len(matr)-1;i++{
		for j:=0;j<len(matr);j++{
			for k := 1;k<len(matr);k++ {
				if isZeroBorderedSquare(matr, i, j, k){
					fmt.Println(i,j,k)
				}
			}
		}
	}



}

func isZeroBorderedSquare(matrix [][]int,row,col ,squareLen int )bool{
	if col+squareLen >= len(matrix) || row+squareLen >= len(matrix) {
		return false
	}
	for i:= col ;i<=col+squareLen;i++{
		if matrix[row][i]!=0 ||  matrix[row+squareLen][i]!=0{
			return false
		}
	}
	for i:= row ;i<=row+squareLen;i++{
		if  matrix[i][col]!=0 || matrix[i][col+squareLen]!=0{
			return false
		}
	}
	return true
}

