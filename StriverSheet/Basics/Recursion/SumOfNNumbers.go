/*
Sum of first n numbers using recursion
*/

package main

import "fmt"

func sum(n int) int {
	if n <= 0 {
		return 0
	}
	return (n + sum(n-1))
}

func main() {
	fmt.Printf("%d", sum(6))
}