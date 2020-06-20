package main

import (
	"reflect"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ShiftLinkedList(head *LinkedList, k int) *LinkedList {
	var prev *LinkedList
	var endPrev *LinkedList
	temp := head
	temp1 := head
	lenTemp1 := 0
	count := 0
	for temp1 != nil {
		temp1 = temp1.Next
		lenTemp1++
	}
	k = k % lenTemp1
	if k < 0 {
		k = lenTemp1 + k
	}
	for temp != nil {
		if count == k {
			prev = head
		}
		if count > k {
			prev = prev.Next
		}
		count++
		endPrev = temp
		temp = temp.Next
	}
	endPrev.Next = head
	head = prev.Next
	prev.Next = nil
	return head
}
func newLinkedList(n int) *LinkedList { return &LinkedList{Value: n} }

func linkedListToArray(head *LinkedList) []int {
	array := []int{}
	current := head
	for current != nil {
		array = append(array, current.Value)
		current = current.Next
	}
	return array
}
func main() {

	head := newLinkedList(0)
	head.Next = newLinkedList(1)
	head.Next.Next = newLinkedList(2)
	head.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next = newLinkedList(4)
	head.Next.Next.Next.Next.Next = newLinkedList(5)
	result := ShiftLinkedList(head, 2)
	array := linkedListToArray(result)

	expected := []int{4, 5, 0, 1, 2, 3}
	if !reflect.DeepEqual(array, expected) {
		panic("erorr in logi")
	}

}
