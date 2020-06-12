package main

import "fmt"


//Find the k the smalles element from unsorted array

func main() {
	//Using quicksort approach or finding the reuqired parition
	fmt.Println(QuickPartition([]int{ 32, 92, 17, 16, 34}, 3))
}
func QuickPartition(array []int, k int) int {
	return Partition(array, 0, len(array)-1, k-1)
}

func Partition(ar []int, low, high, k int) int {
	for low <= high {
		partition := getPartition(ar, low, high)
		if partition == k {
			return ar[partition]
		}
		if partition > k {
			high = partition - 1
		} else {
			low = partition + 1
		}
	}
	return -1
}

func getPartition(ar []int, low, high int) int {
	part := ar[low]
	l := low
	h := high
	for l < h {
		for l <= high && ar[l] <= part {
			l++
		}

		for h >= 0 && ar[h] > part {
			h--
		}

		if l < h {
			ar[l], ar[h] = ar[h], ar[l]
		}

	}
	ar[low], ar[h] = ar[h], part

	return h
}
