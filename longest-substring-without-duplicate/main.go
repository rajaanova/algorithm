package main

import "fmt"

func main() {
	fmt.Println(LongestSubstringWithoutDuplication("abcdeabcdefc"))
}

func LongestSubstringWithoutDuplication(str string) string {
	lastUpdated := make(map[rune]int)
	last := 0
	maxLen := 0
	lowInd := 0
	highInd := 0
	for i,run := range str {
		val,ok := lastUpdated[run]
		if !ok || val < last{
			if maxLen < i - last {
				lowInd = last
				highInd = i
				maxLen = i - last
			}
		}else{
			last = val+1
		}
		lastUpdated[run]=i
	}
	return str[lowInd:highInd+1]
}

