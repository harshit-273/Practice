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
		if value > freq {
			freq = value
			ch = index
		}
	}
	return
}

func main() {
	freq, ch := MaxFrequency("harshita")
	fmt.Printf("%c is present %d times in the string", ch, freq)
}
