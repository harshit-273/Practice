package main

import "fmt"

/*
	- getting the nth bit fromm the given number
*/

func NthBit(num int, bit int) (outBit int) {
	outBit = 0
	if num&(1<<bit) != 0 {
		outBit = 1
	}
	return
}

func main() {
	var num int = 5
	var bit int = 2
	fmt.Printf("%d is the %dth bit in %d", NthBit(num, bit), bit, num)
}
