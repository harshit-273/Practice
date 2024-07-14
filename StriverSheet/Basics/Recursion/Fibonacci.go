/*
Given an integer N. Print the Fibonacci series up to the Nth term.

Fibonacci series is the series which starts with two terms and as the series progresses the next term would be sum of current and previous number.
e.g., 0, 1, 1, 2, 3, 5, 8, 13...
*/

package main

import "fmt"

func fibonacci(term int) (series []int) {
	if term == 1 {
		series = append(series, 0)
	} else if term == 2 {
		series = append(series, 0)
		series = append(series, 1)
	} else {
		seriesMinus1 := fibonacci(term-1)
		series = append(seriesMinus1, (seriesMinus1[term-2] + seriesMinus1[term-3]) )
	}
	return series
}

func main() {
	fmt.Printf("%v", fibonacci(10))
}