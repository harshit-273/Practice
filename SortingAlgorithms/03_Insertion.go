package main

import (
	"fmt"
)

func main() {
	var a []int = []int{0, 1, 2, 3, 4, 5}
	var temp, temp_ind int
	for i:=0; i<len(a); i++ {
		temp = 0
		temp_ind = 0
		for j:=0; j<=i; j++ {
			if a[i]<=a[j] {
				temp = a[i]
				temp_ind = j
				break
			}
		}
		for k:=i; k>temp_ind; k-- {
			a[k] = a[k-1]
		}
		a[temp_ind]= temp
	}
	fmt.Println(a)
}