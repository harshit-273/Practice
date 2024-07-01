/*
Problem Statement: Given an integer N, return true if it is a palindrome else return false.

A palindrome is a number that reads the same backward as forward. For example, 121, 1331, and 4554 are palindromes because they remain the same when their digits are reversed.

We would be reversing a number and comparing if it is same as the original number.
*/

package main

import "fmt"

func main() {
	var num int = 34543
	var reversedNum int = 0
	var newNum int = num
	for num>0 {
		reversedNum = (reversedNum*10) + (num%10)
		num = num/10
	}
	
	if reversedNum == newNum {
		fmt.Printf("%d is a Palindrome", newNum)
	} else {
		fmt.Printf("%d is not a Palindrome", newNum)
	}
}