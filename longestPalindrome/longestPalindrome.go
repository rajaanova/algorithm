package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(LongestPalindromicSubstring("abaxyzzyxf"))
}
func LongestPalindromicSubstring(str string) string {
	maxLen := math.MinInt32
	minIndex := 0
	maxIndex := 0
	for i:=0;i<len(str);i++{
		for j:=i+1;j<=len(str);j++{
			//fmt.Println(str[i:j])
			if IsPalin(str[i:j]) && j-i > maxLen {
				maxLen = j - i
				minIndex = i
				maxIndex = j
			}
		}
	}
	return str[minIndex:maxIndex]
}


func IsPalin(str string) bool  {
	for i,j:=0,len(str)-1;i<j;i,j =i+1,j-1{
		if str[i] != str[j] {
			return false
		}
	}
	return true

}
