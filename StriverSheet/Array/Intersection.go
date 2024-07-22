/*
Given two sorted arrays, arr1, and arr2 of size n and m. Find the intersection of two sorted arrays.

The intersection of two arrays can be defined as only the common elements in the two arrays.

Here, we would b using 2pointers.
*/

package main

import "fmt"

func main() {
	var arr1 []int = []int{1, 2, 3, 4, 4, 4, 6, 7, 8}
	var arr2 []int = []int{4, 4, 5, 6, 9}
	
	fmt.Println("First array:", arr1, "Seconds array:", arr2)
	
	var l1 int = len(arr1)
	var l2 int = len(arr2)
	
	var intersectionWithDup []int = make([]int, 0)
	var i int = 0
	var j int = 0
	
	for i<l1 && j<l2 {
		if arr1[i] == arr2[j] {
			intersectionWithDup = append(intersectionWithDup, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	fmt.Println("Intersection with duplicates:", intersectionWithDup)
	
	var intersectionWithoutDup []int = make([]int, 0)
	var isPresent map[int]int = make(map[int]int)

	var p int = 0
	var q int = 0
	
	for p<l1 && q<l2 {
		if arr1[p] == arr2[q] {
			_, isThere := isPresent[arr1[p]]
			if !isThere {
				intersectionWithoutDup = append(intersectionWithoutDup, arr1[p])
				isPresent[arr1[p]] = 1
			}
			p++
			q++
		} else if arr1[p] < arr2[q] {
			p++
		} else {
			q++
		}
	}
	fmt.Println("Intersection without duplicates:", intersectionWithoutDup)
}