package main

import (
	"reflect"
	"testing"
)

func TestMaxSumIncreasingSubsequence(t *testing.T) {
	maxSum , arraySeq := MaxSumIncreasingSubsequence([]int{10, 70, 20, 30, 50, 11, 30})
	if maxSum != 110{
		t.Errorf("error getting max sum %d",maxSum)
	}
	if !reflect.DeepEqual(arraySeq,[]int{10,20,30,50}){
		t.Errorf("array not equal %v",arraySeq)
	}
}