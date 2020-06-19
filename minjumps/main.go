package main

import "fmt"

func main() {
	fmt.Println(MinNumberOfJumps([]int{1,0,1}))
	fmt.Println(MinNumberOfJumps([]int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3}))
}

func MinNumberOfJumps(a []int) int {
	arrayVal := make([]int,len(a))
	for i, _ := range arrayVal{
		arrayVal[i] = -1
	}
	arrayVal[0] = 0
	for i:=1 ;i<len(a);i++{
		for j:=0;j<i;j++{
			if i-j <= a[j] && (arrayVal[i] == -1 || arrayVal[i] > arrayVal[j] + 1) {
				arrayVal[i] =  arrayVal[j] + 1
			}
		}
	}
	return arrayVal[len(a)-1]



}
