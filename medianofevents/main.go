package main

import (
	"container/heap"
	"fmt"
)

type arrayInt []int

func (a arrayInt) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a arrayInt) Swap(i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
func (a *arrayInt) Top() interface{} {
	return (*a)[0]
}

func (a *arrayInt) Push(x interface{}) {
	va := x.(int)
	*a = append(*a, va)
}

func (a *arrayInt) Pop() interface{} {
	t := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return t
}

func (a arrayInt) Len() int {
	return len(a)
}

type arrayInDecresaing []int

func (a arrayInDecresaing) Less(i, j int) bool {
	return a[i] > a[j]
}
func (a *arrayInDecresaing) Top() interface{} {
	return (*a)[0]
}
func (a arrayInDecresaing) Swap(i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func (a *arrayInDecresaing) Push(x interface{}) {
	va := x.(int)
	*a = append(*a, va)
}

func (a *arrayInDecresaing) Pop() interface{} {
	t := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return t
}

func (a arrayInDecresaing) Len() int {
	return len(a)
}

func main() {
	checking := []int{ 8,7,6,5,4,3,2,1,0}
	hc := arrayInt(checking)
	heap.Init(&hc)
	mediaSequence := []int{1, 0, 3, 5, 2, 0, 1}
	//sort.Ints(mediaSequence)
	a := &arrayInt{}
	b := &arrayInDecresaing{}

	for i := 0; i < len(mediaSequence); i++ {
		fmt.Println(getMedian(mediaSequence[i], a, b))
	}

}

func getMedian(i int, increasing *arrayInt, decreasing *arrayInDecresaing) float64 {

	if len(*increasing) == 0 {
		heap.Push(increasing, i)
	} else {
		if i >= increasing.Top().(int) {
			heap.Push(increasing, i)
		} else {
			heap.Push(decreasing, i)
		}
	}

	if len(*increasing) > len(*decreasing)+1 {
		heap.Push(decreasing, heap.Pop(increasing))
	} else if decreasing.Len() > increasing.Len() {
		heap.Push(increasing,heap.Pop(decreasing))
	}
	if increasing.Len() == decreasing.Len() {
		return 0.5 * float64(((increasing.Top().(int)) + (decreasing.Top().(int))))
	} else {
		return float64(increasing.Top().(int))
	}
}
