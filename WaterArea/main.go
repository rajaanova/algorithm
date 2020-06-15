package main

import "fmt"

func main() {
	output := WaterArea([]int{0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3})
	if output != 48 {
		fmt.Println("error scenario")
	}
}
func WaterArea(heights []int) int {
	leftMax := -1
	leftMaxArea := make([]int, len(heights))
	for i, leftMaxVal := range heights {
		MaxVal := max(leftMaxVal, leftMax)
		if MaxVal > leftMax {
			leftMax = MaxVal
		}
		leftMaxArea[i] = leftMax
	}
	rightMax := -1
	rightMaxArea := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		MaxVal := max(heights[i], rightMax)
		if MaxVal > rightMax {
			rightMax = MaxVal
		}
		rightMaxArea[i] = rightMax
	}
	sum := 0
	for i, val := range heights {
		minArea := min(leftMaxArea[i], rightMaxArea[i])
		if minArea > val {
			sum = sum + minArea - val
		}
	}

	return sum

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
