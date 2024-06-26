/*
Given an integer N return the reverse of the given number

We would be extracting the digits one by one and multiplying the digit to 10 and adding the next extracted digit to it continously till the original number is not zer
*/

package main

import "fmt"

func main() {
	var num int = 34504
	var newNum int = 0
	for num>0 {
		newNum = (newNum*10) + (num%10)
		num = num/10
	}
	fmt.Println(newNum)
}