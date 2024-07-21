/*
Given an array of N integers, left rotate the array by one place.

We would be using one tem variable to store one value here and at last after moving all the elements to left, replace the last element with temp.
*/

package main

import "fmt"

func main() {
	var arr []int = []int{2, 4, 5, 8, 2, 5, 7, 6}
	
	fmt.Println("Array before left rotate by 1:", arr)
	
	var l int = len(arr)
	var temp int = arr[0]
	for i:=1; i<l; i++ {
		arr[i-1] = arr[i]
	}
	arr[l-1] = temp
	
	fmt.Println("Array after left rotate by 1:", arr)
}