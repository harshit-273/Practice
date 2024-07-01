/*
Given two integers N1 and N2, find their greatest common divisor.

The Greatest Common Divisor of any two integers is the largest number that divides both integers.

We would be using a Euclidian principle here:
gcd(num1, num2) = gcd(num1-num2, num2) if num1>num2

And to further make it more efficient we would be using:
gcd(num1%num2) where num1>num2
*/

package main

import "fmt"

func main() {
	var num1, num2 int = 45, 9
	var tmp int
	if num1 > num2 {
		tmp = num1
		num1 = num2
		num2 = tmp
	}
	for (num1 > 0) && (num2 > 0) {
		tmp = num2%num1
		num2 = num1
		num1 = tmp
	}
	fmt.Printf("%d is the GCD", num2)
}