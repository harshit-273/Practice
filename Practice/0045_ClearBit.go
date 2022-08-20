package main

import "fmt"

/*
	- clearing the bit for the given number.
*/

func ClearBit(num int, bit int) (outNum int) {
	outNum = num & ((1 << bit) ^ (-1))
	return
}

func main() {
	var bit int = 3
	var num int = 7
	fmt.Printf("%dth bit is cleared for the number %d, which resulted in %d", bit, num, ClearBit(num, bit))
}
