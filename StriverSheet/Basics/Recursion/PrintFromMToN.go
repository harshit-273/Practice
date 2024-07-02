/*
Printing m to n using recursion
*/

package main

import "fmt"

func printMN(m int, n int) {
	if m>n {
		return
	}
	fmt.Printf("%3d", m)
	m++
	printMN(m, n)
}
func main() {
	printMN(5, 8)
}