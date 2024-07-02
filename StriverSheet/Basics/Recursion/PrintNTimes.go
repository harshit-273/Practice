/*
Understanding recursion by printing N times. 

Recursion is when a function calls itself again and again till a condition is met.
*/

package main

import "fmt"

var count int = 5

func recExample() {
	if count<=0 {
		return
	}
	fmt.Println("Harshit")
	count--
	recExample()
}

func main() {
	recExample()
}