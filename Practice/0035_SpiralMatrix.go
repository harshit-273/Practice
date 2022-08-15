package main

import "fmt"

/*
	- Here we are printing the matrix in spiral order
*/

func SpiralMatrix(arr [][]int, rows int, cols int) {
	var rowStart, rowEnd, colStart, colEnd int = 0, rows - 1, 0, cols - 1

	for rowStart <= rowEnd && colStart <= colEnd {
		// row start
		for col := colStart; col <= colEnd; col++ {
			fmt.Printf("%4d", arr[rowStart][col])
		}
		rowStart += 1

		// column end
		for row := rowStart; row <= rowEnd; row++ {
			fmt.Printf("%4d", arr[row][colEnd])
		}
		colEnd -= 1

		// row end
		for col := colEnd; col >= colStart; col-- {
			fmt.Printf("%4d", arr[rowEnd][col])
		}
		rowEnd -= 1

		// column start
		for row := rowEnd; row >= rowStart; row-- {
			fmt.Printf("%4d", arr[row][colStart])
		}
		colStart += 1
	}
}

func main() {
	var Array2d [][]int = make([][]int, 0)
	var input int
	fmt.Println("Please the enter the 2d array of size 6x4:")
	for i := 0; i < 6; i++ {
		Array1d := make([]int, 0)
		for j := 0; j < 4; j++ {
			fmt.Scan(&input)
			Array1d = append(Array1d, input)
		}
		Array2d = append(Array2d, Array1d)
	}

	SpiralMatrix(Array2d, len(Array2d), len(Array2d[0]))
}
