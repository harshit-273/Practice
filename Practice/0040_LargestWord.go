package main

import (
	"fmt"
	"math"
)

/*
	- Here we are finding the word with the largest length in the sentence.
*/

func LargestWord(str string) (outStr string) {
	var start int = -1
	var end int = -1
	var ptr int = 0
	var largeWordLength int = math.MinInt
	var currLen int = 0
	outStr = "not present"

	for ptr < len(str) {
		if str[ptr] != ' ' {
			currLen++
		} else {
			currLen = 0
		}
		if currLen > largeWordLength {
			largeWordLength = currLen
			start = ptr - (largeWordLength - 1)
			end = ptr
		}
		ptr++
	}

	if start != -1 && end != -1 {
		outStr = str[start:(end + 1)]
	}

	return
}

func main() {
	fmt.Printf("Largest word is %q", LargestWord(""))
}
