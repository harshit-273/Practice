/*
Given an integer N, return true it is an Armstrong number otherwise return false.

An Amrstrong number is a number that is equal to the sum of its own digits each raised to the power of the number of digits.

We would be extracting each digits > cubing them > adding them
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var storedNum int = 1634
	var num int = storedNum
	var sum = 0
	var sli []int = make([]int, 0)
	
	for num >0 {
		sli = append(sli, num%10)
		num = num/10
	}
	var po float64 = float64(len(sli))
	for _, v := range sli {
		sum = sum + int(math.Pow(float64(v), po))
	}
	
	if sum == storedNum {
		fmt.Printf("%d is a Amrstrong number", storedNum)
	} else {
		fmt.Printf("%d is not a Amrstrong number", storedNum) 
	}
}