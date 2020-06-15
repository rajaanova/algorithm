package main

import "fmt"

func main() {
	fmt.Println(NumberOfWayForChange([]int{1, 5}, 6))
}
func NumberOfWayForChange(denom []int, k int) int {
	array := make([]int, k+1)
	array[0] = 1
	for _, deno := range denom {
		for i := 1; i < len(array); i++ {
			if i < deno {
				continue
			}
			array[i] = array[i] + array[i-deno]
		}
	}
	return array[k]
}
