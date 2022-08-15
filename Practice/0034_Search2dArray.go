package main

import "fmt"

/*
	- Here we are searching an element in the 2D-Array
*/

func Search2dArray(arr [][]int, key int) (int, int) {
	for outerIndex, innerArray := range arr {
		for innerIndex, value := range innerArray {
			if value == key {
				return outerIndex, innerIndex
			}
		}
	}
	return -1, -1
}

func main() {
	var Array2d [][]int = make([][]int, 0)
	var input int
	fmt.Println("Please the enter the 2d array of size 4x3:")
	for i := 0; i < 4; i++ {
		Array1d := make([]int, 0)
		for j := 0; j < 3; j++ {
			fmt.Scan(&input)
			Array1d = append(Array1d, input)
		}
		Array2d = append(Array2d, Array1d)
	}

	var searchKey int
	fmt.Print("Please enter the number that needs to be searched in the 2d Array: ")
	fmt.Scan(&searchKey)

	outerIn, innerIn := Search2dArray(Array2d, searchKey)

	if outerIn != -1 && innerIn != -1 {
		fmt.Printf("The search key %d can be found at [%d][%d]", searchKey, outerIn, innerIn)
	} else {
		fmt.Println("Search key not found in the array")
	}
}
