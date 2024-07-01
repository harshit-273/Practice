/*
Given an integer N, check whether it is prime or not. 

A prime number is a number that is only divisible by 1 and itself and the total number of divisors is 2.

We would be checking the number till it squareroot as it would be the largest number after which the divisors would repeat again
*/

package main

import "fmt"

func main() {
	var num int = 69
	var count int = 0
	
	for i:=1; i*i<=num; i++ {
		if num%i == 0 {
			count++
			if i*i != num {
				count++
			}
		}
	}
	
	if count == 2 {
		fmt.Printf("%d is a prime number", num)
	} else {
		fmt.Printf("%d is not a prime number", num)
	}
	
}