/*
Given an integer N, return all divisors of N.

A divisor of an integer N is a positive integer that divides N without leaving a remainder. In other words, if N is divisible by another integer without any remainder, then that integer is considered a divisor of N.

We would be printing all the divisors by going till squareroot of the number
*/

package main 

import "fmt"

func main() {
	var num int = 69
	var divisors []int = make([]int, 0)
	for i:=1; i*i<=num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
			if i*i != num {
				divisors = append(divisors, num/i)
			}
		}
	}
	
	fmt.Println(divisors)
}