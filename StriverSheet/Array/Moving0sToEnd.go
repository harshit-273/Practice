/*
You are given an array of integers, your task is to move all the zeros in the array to the end of the array and move non-negative integers to the front by maintaining their order.

We would be replacing the non-zero values to front using 2pointers.
*/

package main

import "fmt"

func main() {
	var arr []int = []int{0, 2, 4, 0, 5, 8, 0, 0, 2, 5, 7, 6, 0}
	fmt.Println("Before the array being processed to move the zero to end:", arr)
	
	/* For a usecase where there is no zero in the array. */
	var j int = -1
	
	/* For finding 1st 0 */
	var l int = len(arr)
	for i:=0; i<l; i++ {
		if arr[i] == 0 {
			j = i
			break
		}
	}
	if j == -1 {
		fmt.Println("Same array as there are no zeros:", arr)
		return
	}
	
	for k:=(j+1); k<l; k++ {
		if arr[k] != 0 {
			arr[j] = arr[k]
			arr[k] = 0
			j++
		}
	}
	fmt.Println("After the array being processed to move the zero to end:", arr)
}