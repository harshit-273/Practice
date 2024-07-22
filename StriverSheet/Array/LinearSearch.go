/*
Given an array, and an element num the task is to find if num is present in the given array or not. If present print the index of the element or print -1.

For first occurrance, break the loop after finding it.
For Last occurrance, iterate the entire loop.
For all occurrance, store in an array.
*/

package main

import "fmt"

func main() {
	var arr []int = []int{0, 2, 4, 0, 5, 8, 0, 0, 2, 5, 7, 6, 0}
	
	var search map[int]([]int) = make(map[int]([]int))
	
	for ind, val := range arr {
		search[val] = append(search[val], ind)
	}
	
	fmt.Println("All occurrance of all the values and with their location for  array -", arr,"\n=>", search)
}