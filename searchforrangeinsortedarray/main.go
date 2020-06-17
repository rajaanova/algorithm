package main

import "fmt"

func main() {
	fmt.Println(SearchForRange([]int{5, 7, 7, 8, 8, 10}, 10))
}
func SearchForRange(array []int, target int) []int {
	low := 0
	high := len(array)-1
	for low <= high {
		mid := low + (high-low)/2
		if array[mid] == target{
			lowerSide := mid -1
			higherSide := mid + 1
			for lowerSide >= 0 && array[lowerSide] == target {
				lowerSide -= 1
			}
			for higherSide <= len(array)-1 && array[higherSide] == target {
				higherSide += 1
			}

			return []int{lowerSide+1,higherSide-1}
		}

		if array[mid] > target {
			high = mid -1
		}else {
			low = mid + 1
		}

	}

	return []int{-1,-1}
}
