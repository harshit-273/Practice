/*
Given an integer N and an array of size N-1 containing N-1 numbers between 1 to N. Find the number(between 1 to N), that is not present in the given array.

Here we would be using few different methods to solve this problem.
1) checking if every number is present in the array one after another.

2) Using Hashing: marking the numbers which are present and then checking all the numbers one by one.

3) Using summation:
adding all numbers from 1 to N and subtracting addition of numbers in the array from it

4) Using XOR:
XOR of any two numbers is ((a AND b) OR (a OR b))
*/

package main 

import "fmt"

func main() {
	var N int = 5
	var n int = N-1
	arr := [4]int{1, 2, 3, 5}
	
	// Summation method
	var sum int = (N*(N+1))/2
	var arrSum int = 0
	for _, val := range arr {
		arrSum = arrSum + val
	}
	
	fmt.Println("Missing number is", sum-arrSum)
	
	// XOR method
	var xor1toN int = 0
	var xorOfArr int = 0
	for i:=0; i<n; i++ {
		xorOfArr = xorOfArr^arr[i] // XOR of all the numbers in array
		xor1toN = xor1toN^(i+1) // XOR of all numbers from 1 to (N-1)
	}
	xor1toN = xor1toN^N // XOR with N
	
	fmt.Println("Missing number is", xor1toN^xorOfArr)
}