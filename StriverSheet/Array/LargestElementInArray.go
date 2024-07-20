/*
Given an array, we have to find the largest element in the array.

We would be assuming the 1st element to be the largest and later on updating the largest value as traverse through the array.
*/

package main

import "fmt"

func main() {
	var arr []int = []int{4, 3, 7, 4, 9, 2, 5}
	var max = arr[0]
	
	for _, val := range arr[1:] {
		if max < val {
			max = val
		}
	}
	
	fmt.Print("Largest number in array ", arr, " is ", max)
}