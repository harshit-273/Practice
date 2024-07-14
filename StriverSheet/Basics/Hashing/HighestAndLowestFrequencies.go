/*
Given an array of size N. Find the highest and lowest frequency element.

We would be using a Map to store the frequency and later iterating through it to get the highest and lowest frequencies
*/

package main

import "fmt"

func main() {
	var arr []int = []int{45, 54, 34, 32, 12, 56, 67, 32, 12, 54, 98}
	var intMap map[int]int = make(map[int]int)
	
	for _, val := range arr {
		intMap[val]++
	}
	
	var minFreq int = len(arr)
	var maxFreq int = 0
	var minNum int = arr[0]
	var maxNum int = arr[0]
	for ind, val := range intMap {
		if val <= minFreq {
			minFreq = val
			minNum = ind
		}
		if val >= maxFreq {
			maxFreq = val
			maxNum = ind
		}
	}
	
	fmt.Println("Number with highest frequency of", maxFreq, "is", maxNum)
	fmt.Println("Number with lowest frequency of", minFreq, "is", minNum)
}