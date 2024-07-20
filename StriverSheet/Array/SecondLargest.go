/*
Given an array, find the second smallest and second largest element in the array. Print ‘-1’ in the event that either of them doesn’t exist.
*/

package main

import "fmt"

func main() {
	var arr []int = []int{2, 4, 5, 8, 2, 5, 7, 6}
	var lar int = arr[0]
	var secondLar int = -1
	
	for _, val := range arr[1:] {
		if val > lar {
			secondLar = lar
			lar = val
		} else if val > secondLar {
			secondLar = val
		}
	}
	
	fmt.Println("Largest in the array -", arr, "is", lar, "and 2nd Largest is", secondLar)
}