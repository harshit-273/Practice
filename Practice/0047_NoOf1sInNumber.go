package main

import "fmt"

/*
	- we will be counting the number of 1's present in the binary form of a number
*/

func NoOf1sInNumber(num int) (count int) {
	count = 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return
}

func main() {
	var num int = 0
	fmt.Printf("%d is the number of 1's present in the number %d", NoOf1sInNumber(num), num)
}
