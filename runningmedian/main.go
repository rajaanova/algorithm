package main

import (
	"container/heap"
	"fmt"
	"math"
	"runtime"
)

type arr []int

func (a *arr) Len() int {
	return len(*a)
}

func (a *arr) Less(i, j int) bool {
	return (*a)[i] < (*a)[j]
}

func (a arr) Swap(i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func (a *arr) Push(x interface{}) {
	*a = append(*a, x.(int))
}

func (a *arr) Pop() interface{} {
	od := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return od

}

type maxheapArr []int

func (a maxheapArr) Len() int {
	return len(a)
}

func (a maxheapArr) Less(i, j int) bool {
	return (a)[i] > (a)[j]
}

func (a maxheapArr) Swap(i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func (a *maxheapArr) Push(x interface{}) {
	*a = append(*a, x.(int))
}

func (a *maxheapArr) Pop() interface{} {
	od := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return od

}

func main() {
	fmt.Println(runtime.NumCPU())
	upperHalf := &arr{}
	lowerHalf := &maxheapArr{}
	heap.Init(upperHalf)
	heap.Init(lowerHalf)
	stream := []int{23, 34, 54, 12, 32, 3, 87, 57}
	medians := make([]float64, len(stream))
	for i, val := range stream {
		medians[i] = getMedian(val, upperHalf, lowerHalf)
	}
	fmt.Println("jjk")
	fmt.Println(medians)
}
func getMedian(i int, upperhalf *arr, lowerhalf *maxheapArr) float64 {
	//isnert
	if len(*lowerhalf) == 0 {
		heap.Push(lowerhalf, i)
		return float64(i)
	}

	x := heap.Pop(lowerhalf)
	if x.(int) > i {
		heap.Push(lowerhalf, i)
	} else {
		heap.Push(upperhalf, i)
	}
	heap.Push(lowerhalf, x)

	if math.Abs(float64(len(*upperhalf)-len(*lowerhalf))) >= 2 {
		if len(*upperhalf) > len(*lowerhalf) {
			heap.Push(lowerhalf, heap.Pop(upperhalf))
		} else {
			heap.Push(upperhalf, heap.Pop(lowerhalf))
		}
	}
	y := heap.Pop(lowerhalf)
	z := heap.Pop(upperhalf)
	heap.Push(upperhalf, z)
	heap.Push(lowerhalf, y)
	if len(*lowerhalf) == len(*upperhalf) {
		return (float64(y.(int)) + float64(z.(int))) / 2
	}
	if len(*lowerhalf) > len(*upperhalf) {
		return float64(y.(int))
	} else {
		return float64(z.(int))
	}

}
