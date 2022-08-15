package main

import "fmt"

/*
	- Here we are transposing a matrix
*/

func TransposeMatrix(arr [][]int) (ar [][]int) {
	ar = make([][]int, len(arr[0]))
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(arr); j++ {
			ar[i] = append(ar[i], arr[j][i])
		}
	}
	return
}

func main() {
	fmt.Println(TransposeMatrix([][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11}}))
}
