/*
Given an array that contains only 1 and 0 return the count of maximum consecutive ones in the array.

keeping the track of current consecutive 1s and updating the value of max consecutive 1's
*/

package main

import "fmt"

func main() {
	var arr []int = []int{1, 1, 0, 1, 1, 1, 0, 1, 1}
	
	var current int = 0
	var max int = 0
	
	for _, val := range arr {
		if val == 1 {
			current++
			if max < current {
				max = current
			}
		} else {
			current = 0
		}
	}
	
	fmt.Println("Maximum consecutive 1s are", max)
}