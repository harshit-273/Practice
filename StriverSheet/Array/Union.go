/*
Given two sorted arrays, arr1, and arr2 of size n and m. Find the union of two sorted arrays.

The union of two arrays can be defined as the common and distinct elements in the two arrays.

Here, we can have it both ways, allowing same values more than once if already there are more values in one of the set or not allowing same values more than once.

NOTE: Elements in the union should be in ascending order.

We would be using 2pointers, as we already know the arrays are sorted.
*/

package main

import "fmt"

func main() {
	var arr1 []int = []int{1, 2, 3, 4, 4, 4, 6, 7, 8}
	var arr2 []int = []int{4, 4, 5, 6, 9}
	
	fmt.Println("First array:", arr1, "Seconds array:", arr2)
	
	var l1 int = len(arr1)
	var l2 int = len(arr2)
	
	var unionWithDup []int = make([]int, 0)
	var i int = 0
	var j int = 0
	
	// Allowing duplicate values in below.
	for (i<l1) && (j<l2) {
		if arr1[i] == arr2[j] {
			unionWithDup = append(unionWithDup, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			unionWithDup = append(unionWithDup, arr1[i])
			i++
		} else {
			unionWithDup = append(unionWithDup, arr2[j])
			j++
		}
	}
	for i<l1 {
		unionWithDup = append(unionWithDup, arr2[i])
		i++
	}
	for j<l2 {
		unionWithDup = append(unionWithDup, arr2[j])
		j++
	}
	fmt.Println("Union if duplicates are allowed:", unionWithDup)
	
	//Not allowing duplicates
	var unionWithoutDup []int = make([]int, 0)
	var isPresent map[int]int = make(map[int]int)
	var p int = 0
	var q int = 0
	for (p<l1) && (q<l2) {
		if arr1[p] == arr2[q] {
			_, isThere := isPresent[arr1[p]]
			if !isThere {
				unionWithoutDup = append(unionWithoutDup, arr1[p])
				isPresent[arr1[p]] = 1
			}
			p++
			q++
		} else if arr1[p] < arr2[q] {
			_, isThere := isPresent[arr1[p]]
			if !isThere {
				unionWithoutDup = append(unionWithoutDup, arr1[p])
				isPresent[arr1[p]] = 1
			}
			p++
		} else {
			_, isThere := isPresent[arr2[q]]
			if !isThere {
				unionWithoutDup = append(unionWithoutDup, arr2[q])
				isPresent[arr2[q]] = 1
			}
			q++
		}
	}
	for p<l1 {
		unionWithoutDup = append(unionWithoutDup, arr1[p])
		p++
	}
	for q<l2 {
		unionWithoutDup = append(unionWithoutDup, arr2[q])
		q++
	}
	
	fmt.Println("If duplicates are not allowed:", unionWithoutDup)
}