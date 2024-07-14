/*
Given an array, we have found the number of occurrences of each element in the array.

We would be using Maps here. Which can store a key and can have value.
*/

package main

import "fmt"

func main() {
	var intMap map[int]int = make(map[int]int)
	var ar []int = []int{45, 54, 34, 32, 12, 56, 67, 32, 12, 54, 98}
	
	for _, val := range ar {
		intMap[val]++
	}
	
	fmt.Printf("%v", intMap)
}