/*
You are given an array. The task is to reverse the array and print it. 

We would be swapping 1st and last indices till the mid.
*/

package main

import "fmt"

func revArr(arr []int, nth int, lastNth int) ([]int) {
	var temp int
	if nth < lastNth {
		temp = arr[nth]
		arr[nth] = arr[lastNth]
		arr[lastNth] = temp
		arr = revArr(arr, nth+1, lastNth-1)
	}
	return arr
}

func main() {
	var ar []int = []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%v", revArr(ar, 0, len(ar)-1))
}