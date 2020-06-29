package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MaxSumIncreasingSubsequence([]int{10, 60, 20, -5, 50, 11, 30}))
}
func MaxSumIncreasingSubsequence(array []int) (int,[]int) {
	sumArray := make([]int,len(array))
	seqArray := make([][]int,len(array))
	sumArray[0] = array[0]
	seqArray[0] = []int{array[0]}

	for i := 1;i<len(array);i++{
		maxTillNow := math.MinInt32
		var indices []int
		for j:=0;j<i;j++{
			if array[i]>array[j] && sumArray[j] > maxTillNow {
				maxTillNow = sumArray[j]
				indices = seqArray[j]
			}
		}
		if len(indices) > 0 {
			sumArray[i] = maxTillNow + array[i]
			seqArray[i] = append(indices,array[i])
		}else{
			sumArray[i] = array[i]
			seqArray[i] = []int{array[i]}
		}

	}

	max := 0
	for i:=1;i<len(array);i++{
		if sumArray[i] > sumArray[max] {
			max = i
		}
	}
	fmt.Println(sumArray)
	return sumArray[max],  seqArray[max]

}
