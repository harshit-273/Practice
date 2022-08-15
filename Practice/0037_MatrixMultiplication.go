package main

import "fmt"

/*
	- Here we are Multiplying the two matrices

	|1 2 3|		|10 11 12 13|
	|4 5 6|		|20 21 23 24|
				|30 31 32 33|
*/

func MatrixMultiplication(m1 [][]int, m2 [][]int) (ansMat [][]int) {
	rowsM1 := len(m1)
	rowsM2 := len(m2)
	colsM1 := len(m1[0])
	colsM2 := len(m2[0])
	if colsM1 != rowsM2 {
		ansMat = [][]int{{-999}}
	} else {
		ansMat = make([][]int, 0)
		for k := 0; k < rowsM1; k++ {
			rowM3 := make([]int, 0)
			for i := 0; i < colsM2; i++ {
				cell := 0
				for j := 0; j < colsM1; j++ {
					cell += m1[k][j] * m2[j][i]
				}
				rowM3 = append(rowM3, cell)
			}
			ansMat = append(ansMat, rowM3)
		}
	}
	return
}

func main() {
	fmt.Println(MatrixMultiplication([][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{10, 11, 12, 13}, {20, 21, 22, 23}, {30, 31, 32, 33}}))
}
