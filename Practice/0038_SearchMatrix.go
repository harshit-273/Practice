package main

import "fmt"

/*
	- Search the given value in matrix, if numbers sorted in ascending for rows nad ascending for columns
*/

func SearchMatrix(mat [][]int, searchKey int) (outerIndex int, innerIndex int) {
	innerIndex = len(mat[0]) - 1
	outerIndex = 0
	for innerIndex >= 0 && outerIndex <= (len(mat)-1) {
		if searchKey == mat[outerIndex][innerIndex] {
			return
		} else if searchKey > mat[outerIndex][innerIndex] {
			outerIndex++
		} else {
			innerIndex--
		}
	}

	innerIndex = -1
	outerIndex = -1
	return
}
func main() {
	fmt.Println(SearchMatrix([][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}, {10, 13, 14}}, 6))
}
