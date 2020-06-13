package main

import "fmt"

func main() {
	points := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	}
	fmt.Println(IsStraightLine(points))
}

func IsStraightLine(points [][]int)  bool {
	if len(points) <=2 {
		return true
	}
	slopeTwoPoints := slope(points[0],points[1])
	for i:=1;i<len(points)-1;i++{
		if slope(points[i],points[i+1]) != slopeTwoPoints {
			return false
		}
	}
	return true
}


func slope(point1, point2 []int ) int {
	y2 := point2[1]
	y1 := point1[1]
	x2 := point2[0]
	x1 := point1[0]
	return (y2-y1)/(x2-x1)
}
