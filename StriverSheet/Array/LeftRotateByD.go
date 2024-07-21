/*
Given an array of integers, rotating array of elements by k elements either left or right.

We would be using a temp array and storing the values which are need to be replaced in the beginning of the array. Then moving remaining elements to left and at last replacing the elements at the end from the temp array.
*/

package main

import "fmt"

func main() {
	var arr []int = []int{2, 4, 5, 8, 2, 5, 7, 6}
	var leftBy int = 10
	
	fmt.Println("Array before Left rotating by", leftBy, "is", arr)
	var l int = len(arr)
	
	/* Performing the below change because the if rotating by same as the size of array results in same array itself */
	leftBy = leftBy%l
	
	var temp []int
	
	/* Using below loop to insert values into temp as using "temp []int = arr[:leftBy]" causes the replace by address(pointer), so if the value in arr changes from here then the values in temp will also be changed */
	for k:=0; k<leftBy; k++ {
		temp = append(temp, arr[k])
	}
	
	for i:=leftBy; i<l; i++ {
		arr[i-leftBy] = arr[i]
	}
	for j:=(l-leftBy); j<l; j++ {
		arr[j] = temp[j-(l-leftBy)]
	}
	fmt.Println("Array after Left rotating by", leftBy, "is", arr)
} 