package main

import "fmt"

func CheckPalindrome(str string) (isPalindrome bool) {
	isPalindrome = false
	var start int = 0
	var end int = len(str) - 1
	for end > start {
		if str[start] != str[end] {
			return
		}
		start++
		end--
	}
	isPalindrome = true
	return
}

func main() {
	var str string = "rac e car"
	var palin bool = CheckPalindrome(str)
	if palin {
		fmt.Printf("%q is a palindrome", str)
	} else {
		fmt.Printf("%q is not a palindrome", str)
	}
	fmt.Println()
}
