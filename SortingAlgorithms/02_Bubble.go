package main

import(
	"fmt"
)

func main() {
	var arr []int = []int{4, 3, 2, 1, 0, 5}
	var temp int
	for i := 0; i<len(arr); i++ {
		for j:=0; j<len(arr)-i-1; j++ {
			if arr[j] >= arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	
	fmt.Println(arr)
}