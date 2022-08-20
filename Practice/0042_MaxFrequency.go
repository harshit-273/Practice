package main

import "fmt"

/*
	- Here we will be finding the character with the highest frequency
*/

func MaxFrequency(str string) (freq int, ch rune) {
	runeStr := []rune(str)
	freqMap := make(map[rune]int)
	for i := 0; i < len(runeStr); i++ {
		freqMap[runeStr[i]]++
	}
	freq = 0
	ch = ' '
	for index, value := range freqMap {
		if value > freq { // if we use ">=" instead of ">" then last character with most frequency will be the answer
			freq = value
			ch = index
		}
	}
	return
}

func main() {
	freq, ch := MaxFrequency("aaabbccc")
	fmt.Printf("%c is present %d times in the string", ch, freq)
}
