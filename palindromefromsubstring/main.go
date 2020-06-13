package main

import "fmt"

func main() {
	s := "abcda"
	queriesRes := make(map[int]bool)
	queries := [][]int{
		{3, 3, 0},
		{1, 2, 0},
		{0, 3, 1},
		{0, 3, 2},
		{0, 4, 1},
	}
	for i, val := range queries{
		firstIndex := val[0]
		secIndex := val[1]
		numberOfReplacement := val[2]
		queriesRes[i] = getPalin(s[firstIndex:secIndex+1],numberOfReplacement)
	}
	fmt.Println(queriesRes)
}

func getPalin(s string, replacement int) bool {
	if len(s) ==1 {
		return true
	}
	keepCount := make(map[rune]int)
	for _, val := range s {
		if _,ok:= keepCount[val];ok {
			delete(keepCount,val)
		}else{
			keepCount[val] = 1
		}
	}
	numberOfOddCount := len( keepCount )
	return numberOfOddCount/2 <= replacement
}
