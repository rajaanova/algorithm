package main


func main() {
	expected := 2
	output := ShiftedBinarySearch([]int{1,2,3,4}, 3)
	if output != expected {
		panic("error ")
	}
	expected = 0
	output = ShiftedBinarySearch([]int{72, 73, 0, 1, 21, 33, 45, 45, 61, 71},  72)
	if expected != output{
		panic("error")
	}
	expected = -1
	output = ShiftedBinarySearch([]int{72, 73, 0, 1, 21, 33, 45, 45, 61, 71},  69)
	if expected != output{
		panic("error")
	}

}
func ShiftedBinarySearch(array []int, target int) int {
	pivot := getPivot(array)
	if target > pivot && target < array[len(array)-1] {
		return binarySearch(array, pivot, len(array)-1, target)
	}
	return binarySearch(array, 0, pivot-1, target)

}

func binarySearch(array []int, low, high , target int ) int {

	for low <= high {
		mid := low + (high-low)/2
		if array[mid] == target {
			return mid
		}

		if target > array[mid] {
			low = mid +1
		}else{
			high = mid - 1
		}
	}
	return -1
}


func getPivot(array []int) int {
	low := 0
	high := len(array)-1
	if array[low] < array[high] {
		return low
	}

	for low <= high {
		mid := low + (high-low)/2
		if mid+1 <= high && array[mid] > array[mid+1] {
			return mid+1
		}
		if mid-1 >= low && array[mid-1] > array[mid] {
			return mid
		}

		if array[mid] > array[low] {
			low = mid + 1
		}else{
			high = mid - 1
		}

	}
	return low
}