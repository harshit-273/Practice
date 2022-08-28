package main

import "fmt"

/*
	- returning the only number which is not repeated two time
*/

func UniqueNo(ls []int) (uniqueNo int) {
	xor := 0
	for _, value := range ls {
		xor ^= value
	}
	return xor
}

func main() {
	ls := []int{1, 2, 3, 4, 1, 2, 3}
	fmt.Printf("%d is the unique number in %v", UniqueNo(ls), ls)
}
