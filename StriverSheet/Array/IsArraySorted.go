/*
Given an array of size n, write a program to check if the given array is sorted in (ascending / Increasing / Non-decreasing) order or not. If the array is sorted then return True, Else return False.

Note: Two consecutive equal values are considered to be sorted.
*/

package main

import "fmt"

func main() {
	var arr []int = []int{2, 4, 5, 5, 7, 8}
	var isSorted = true
	var l int = len(arr)
	for i:=1; i<l; i++ {
		if arr[i] < arr[i-1]{
			isSorted = false
		}
	}
	
	fmt.Println("The array", arr, "is sorted:", isSorted)
}