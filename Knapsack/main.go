package main
import "fmt"

func main() {
	KnapsackProblem([][]int{{1, 2}, {4, 3}, {5, 6}, {6, 7}}, 10)
}
func KnapsackProblem(items [][]int, capacity int) [][]int {
	lenOfItems := len(items) + 1
	valueArray := make([][]int,lenOfItems)
	for i := 0;i<len(valueArray);i++{
		valueArray[i] = make([]int, capacity+1)
	}


	for i, val := range items {
		index := i + 1
		for j := 0 ; j< capacity+1;j++{
			aboveValue := valueArray[index-1][j]
			if  j >= val[1]  {
				lowerValue := valueArray[index-1][j-val[1]]+val[0]
				valueArray[index][j] =	max(aboveValue,lowerValue)
			}else{
				valueArray[index][j] = aboveValue
			}
		}
	}


	fmt.Println(valueArray)
	//TODO : Backtracking required

	return nil

}
func Backtracking(valueArray [][]int) [][]int {
	return nil
}
func max(a , b int) int{
	if a > b {
		return a
	}
	return b
}
