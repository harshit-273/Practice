package main

import (
	"fmt"
)

func merge(arr []int, begin int, mid int, end int) {
	var leftArraySize = mid-begin+1
	var rightArraySize = end-mid
	
	var leftArray []int = make([]int, leftArraySize)
	var rightArray []int = make([]int, rightArraySize)
	
	for i:=0; i<leftArraySize; i++ {
		leftArray[i] = arr[begin+i]
	}
	for i:=0; i<rightArraySize; i++ {
		rightArray[i] = arr[mid+1+i]
	}
	
	var leftArrayPtr int = 0
	var rightArrayPtr int = 0
	var mergedArrayPtr int = begin
	
	for (leftArrayPtr < leftArraySize) && (rightArrayPtr < rightArraySize) {
		if leftArray[leftArrayPtr] <= rightArray[rightArrayPtr] {
			arr[mergedArrayPtr] = leftArray[leftArrayPtr]
			leftArrayPtr++
		} else {
			arr[mergedArrayPtr] = rightArray[rightArrayPtr]
			rightArrayPtr++
		}
		mergedArrayPtr++
	}
	
	for leftArrayPtr < leftArraySize {
		arr[mergedArrayPtr] = leftArray[leftArrayPtr]
		leftArrayPtr++
		mergedArrayPtr++
	}
	for rightArrayPtr < rightArraySize {
		arr[mergedArrayPtr] = rightArray[rightArrayPtr]
		rightArrayPtr++
		mergedArrayPtr++
	}
}

func mergeSort(arr []int, begin int, end int) {
	if begin >= end {
		return
	}
	var mid int = begin + ((end-begin)/2)
	mergeSort(arr, begin, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, begin, mid, end)
}

func main() {
	var arr []int = []int{5, 4, 3, 6, 2, 8}
	
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}