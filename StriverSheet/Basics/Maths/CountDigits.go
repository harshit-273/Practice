/************************************************************************
Given an integer N, return the number of digits in N.
************************************************************************/

package main

import (
	"fmt"
)

func main() {
	var num int = 0
	var count int = 0
	
	for num>0 {
		fmt.Printf("%3d ", num)
		num = num/10
		count++
	}
	fmt.Printf("%d", count)
}