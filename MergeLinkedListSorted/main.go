package main

import "fmt"

func main() {
	list1 := NewLinkedList(2, 6, 7, 8)
	list2 := NewLinkedList(1, 3, 4, 5, 9, 10)
	output := MergeLinkedLists(list1, list2)
	for output != nil {
		fmt.Println(output.Value)
		output = output.Next
	}
}

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func NewLinkedList(val int, others ...int) *LinkedList {
	ll := &LinkedList{Value: val}
	current := ll
	for _, other := range others {
		current.Next = &LinkedList{Value: other}
		current = current.Next
	}
	return ll
}

func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	if headOne == nil {
		return headTwo
	}
	if headTwo == nil {
		return headOne
	}

	var head *LinkedList
	if headOne.Value < headTwo.Value {
		head = headOne
		head.Next = MergeLinkedLists(head.Next, headTwo)
	} else {
		head = headTwo
		head.Next = MergeLinkedLists(head.Next, headOne)
	}
	return head
}
