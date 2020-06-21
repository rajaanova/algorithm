package main

import "fmt"

//https://www.hackerrank.com/challenges/candies/problem

func main() {
		childrenByrating := []int{2,4,2,6,1,7,8,9,2,1}
		fmt.Println(minCandies(childrenByrating))
}

func minCandies(childrenByRating []int) int {
	//We make an array which compares the candies with left seated childres
	leftArray := make([]int,len(childrenByRating))
	rightArray := make([]int,len(childrenByRating))
	for i,_:= range leftArray{
		leftArray[i] = 1
		rightArray[i] = 1
	}
	for i := 1;i<len(leftArray);i++ {
		if childrenByRating[i]>childrenByRating[i-1]{
			leftArray[i] = leftArray[i-1]+1
		}
		if childrenByRating[len(leftArray)-i-1] > childrenByRating[len(leftArray)-i]{
			rightArray[len(leftArray)-i-1] = rightArray[len(leftArray)-i] +1
		}
	}
	sumOfCandies := 0
	for i:=0;i<len(leftArray);i++{
		sumOfCandies = sumOfCandies + max(leftArray[i],rightArray[i])
	}
	return sumOfCandies
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}