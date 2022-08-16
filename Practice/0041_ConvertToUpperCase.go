package main

import "fmt"

/*
	- Here we will be converting thewhole string into Upper string
*/

func ConvertToUpperCase(str *string) {
	runeStr := []rune(*str)
	for i := 0; i < len(runeStr); i++ {
		if runeStr[i] != ' ' && runeStr[i] >= 'a' && runeStr[i] <= 'z' {
			runeStr[i] += 'A' - 'a'
		}
	}
	*str = string(runeStr)
}

func main() {
	var str string = "Harshit Kalavadia"
	fmt.Printf("%q when converted to upper case is ", str)
	ConvertToUpperCase(&str)
	fmt.Printf("%q", str)
}
