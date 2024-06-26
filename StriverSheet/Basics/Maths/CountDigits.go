/************************************************************************
Given an integer N, return the number of digits in N.

Here, we are counting the digits by dividing the number by 10 repeatedly we get zero. So, each time we divide by 10 is added to the count of digits.
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