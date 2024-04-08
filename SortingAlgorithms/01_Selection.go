package main 

import(
	"fmt"
)

func main() {
	var arr []int = []int{ 4, 3, 2, 1, 0, 5}
	var temp int
	for i := 0; i<len(arr); i++ {
		min := i
		for j:=i; j<len(arr); j++ {
			if arr[min] >= arr[j] {
				min = j
			}
		}
		temp = arr[min]
		arr[min] = arr[i]
		arr[i] = temp
	}
	
	fmt.Println(arr)
}