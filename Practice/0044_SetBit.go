package main

import "fmt"

/*
	- setting the bit for the given number
*/

func SetBit(num int, bit int) (outNum int) {
	outNum = num | (1 << bit)
	return
}

func main() {
	var bit int = 1
	var num int = 5
	fmt.Printf("%dth is set for the number %d, which resulted in the %d", bit, num, SetBit(num, bit))
}
