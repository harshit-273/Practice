/*
Given an integer array sorted in non-decreasing order, remove the duplicates in place such that each unique element appears only once. The relative order of the elements should be kept the same.

If there are k elements after removing the duplicates, then the first k elements of the array should hold the final result. It does not matter what you leave beyond the first k elements.

Note: Return k after placing the final result in the first k slots of the array.

We would be using 2-pointers here incresing ptr1 for every element and ptr2 when different element is found.
*/

package main

import "fmt"

func main() {
	var arr []int = []int{2, 2, 5, 5, 5, 6, 7, 7, 8}
	
	fmt.Println("Sorted Array before replacing for unique elements -", arr)
	var l int = len(arr)
	var j int = 0
	for i:=1; i<l; i++ {
		if arr[j] != arr[i] {
			arr[j+1] = arr[i]
			j++
		}
	}
	
	fmt.Println("Sorted Array after replacing the unique elements -", arr)
}