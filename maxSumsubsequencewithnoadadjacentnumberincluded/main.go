package main

import "fmt"

func main() {
	array := []int{10, 70, 20, 30, 50, 11, 30}
	excl := 0
	var  exclNew  int
	prev := array[0]
	//The logic is keep an eye on max values till now and also compare with last adjacent index sum
	for i := 1 ;i <len(array);i++{
		exclNew = prev
		if excl > prev{
			exclNew = excl
		}
		//This includes last adjacent index value so we dont add  prev value with current index value
		prev = array[i] + excl
		excl = exclNew
	}
	if prev > excl{
		fmt.Println(prev)
		return
	}
	fmt.Println(excl)
}
