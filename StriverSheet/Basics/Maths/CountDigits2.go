/************************************************************************
Given an integer N, return the number of digits in N.

As we are counting the numbers by dividing it by 10, we can use log base 10 here. Example: log base 10 of 10000 = 4. Adding 1 to the result of it will give us the exact number of digits in the given number.
************************************************************************/

package main

import (
	"fmt"
	"math"
)

func main() {
	var num int = 4597
	var count int = int(math.Log10(num)) + 1
	fmt.Printf("%d", count)
}