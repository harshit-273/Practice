/*
Factorial

It is the multiplication of numbers from 1 to n
*/

package main

import "fmt"

func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return (n*fact(n-1))
}

func main() {
	fmt.Printf("%d", fact(5))
}