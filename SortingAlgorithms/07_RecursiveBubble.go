package main

import(
	"fmt"
)

func recursiveBubble(arr []int, ptr int, end int) {
	if ptr >= end {
		return
	}
	var curr int = ptr
	var temp int
	for ptr<end {
		if arr[ptr] > arr[ptr+1] {
			temp = arr[ptr]
			arr[ptr] = arr[ptr+1]
			arr[ptr+1] = temp
		}
		ptr++
	}
	recursiveBubble(arr, curr, end-1)
}

func main() {
	var arr []int = []int{4, 0, 2, 5, 4, 3, 2, 1}
	
	recursiveBubble(arr, 0, len(arr)-1)
	
	fmt.Println(arr)
}