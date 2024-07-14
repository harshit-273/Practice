/*
Given a string, check if the string is palindrome or not.

A string is said to be palindrome if the reverse of the string is the same as the string.

We would be checking the nth and lastNth entry and compare them to see if they are same
*/

package main

import "fmt"

func isPalindrome(s string, nth int, lastNth int) (bool) {
	var isPali bool = true
	if nth < lastNth {
		if s[nth] != s[lastNth] {
			isPali = false
		} else {
			isPali = isPalindrome(s, nth+1, lastNth-1)
		}
	}
	return isPali
}

func main() {
	var s string = "aba"
	fmt.Printf("%t", isPalindrome(s, 0, len(s)-1))
}