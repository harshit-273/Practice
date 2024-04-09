package main

import(
	"fmt"
)

func partition(arr []int, low int, high int) int {
	var pivot int = arr[high]
	
	var allLeftPtr = low-1
	
	var temp int
	for i:=low; i<=high; i++ {
		if arr[i] < pivot {
			allLeftPtr++
		
			temp = arr[allLeftPtr]
			arr[allLeftPtr] = arr[i]
			arr[i] = temp
		}
	}
	
	temp = arr[high]
	arr[high] = arr[allLeftPtr+1]
	arr[allLeftPtr+1] = temp
	
	return (allLeftPtr)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		var pi int = partition(arr, low, high)
		
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	var arr []int = []int{10, 7, 8, 9, 1, 2, 5}
	
	quickSort(arr, 0, len(arr)-1)
	
	fmt.Println(arr)
}