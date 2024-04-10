package main

import(
	"fmt"
)

func recInsertion(arr []int, left int, right int) {
	if right <= left {
		return
	}
	var temp_val int = arr[right]
	var temp_ind int = right
	recInsertion(arr, left, right-1)
	for i:=left; i<right; i++ {
		if arr[i] > temp_val {
			temp_ind = i
			break
		}
	}
	for i:=right; i>temp_ind; i-- {
		arr[i]=arr[i-1]
	}
	arr[temp_ind] = temp_val
	return
}

func main() {
	var arr []int = []int{5, 6, 4}
	
	recInsertion(arr, 0, len(arr)-1)
	
	fmt.Println(arr)
}