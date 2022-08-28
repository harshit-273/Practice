package main

import "fmt"

/*
	- return all the subsets of the set
*/

func AllSubsets(set []int) (allSubset [][]int) {
	allSubset = make([][]int, 0)
	var noOfElements int = len(set)
	var noOfSubset int = 1 << noOfElements
	for i := 0; i < noOfSubset; i++ {
		subset := make([]int, 0)
		for j := 0; j < noOfElements; j++ {
			if (i & (1 << j)) != 0 {
				subset = append(subset, set[j])
			}
		}
		allSubset = append(allSubset, subset)
	}
	return
}

func main() {
	set := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v has these subsets - %v", set, AllSubsets(set))
}
