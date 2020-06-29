package main

func main() {
LongestIncreasingSubsequence([]int{5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35})


}




func LongestIncreasingSubsequence(input []int) []int {
	sequence := make([][]int,len(input))
	length := make([]int,len(input))
	sequence[0] = []int{input[0]}
	length[0] = 1
	for i:=1;i<len(input);i++ {
		max := 0
		isZeroIncluded := false

		for j := 0;j<i;j++ {

			if input[i]>input[j] {
				isZeroIncluded = true
				if len(sequence[j]) >= len(sequence[max]) {
					max = j
				}
			}

		}
		if isZeroIncluded{
			length[i] = len(sequence[max])+1
			sequence[i] = append(sequence[max],input[i])
		}else{
			length[i] = 1
			sequence[i]= []int{input[i]}
		}

	}
	max := 0
	for i:=1;i<len(input);i++{
		if len(sequence[i]) > len(sequence[max]) {
			max = i
		}
	}

	return sequence[max]

}