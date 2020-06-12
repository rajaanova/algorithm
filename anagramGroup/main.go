package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"yo", "act", "flop", "tac", "cat", "oy", "olfp"}
	fmt.Println(GroupAnagrams(words))
}
func GroupAnagrams(words []string) [][]string {
	s := make(map[string][]string)
	for _,val := range words {
		vv := sortCharactersInWord(val)
		//st := strings.Split(val,"")
		//sort.Strings(st)
		//vv := strings.Join(st,"")
		s[vv] = append(s[vv], val)
	}
	sss := make([][]string,0)
	for _,v := range s {
		sss = append(sss,v)
	}
	return sss
}
func sortCharactersInWord(s string)string {
	byteSlice := []byte(s)
	sort.Slice(byteSlice, func(i, j int) bool {
		return byteSlice[i] < byteSlice[j]
	})
	return string(byteSlice)
}