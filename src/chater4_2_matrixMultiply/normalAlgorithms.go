package chater4_2_matrixMultiply

import "usefulFunc"

func SquareMatrixMultiply(matA [][]int, matB [][]int) [][]int {
	if !usefulFunc.MatrixesCanMultiply(matA, matB) {
		panic("The two mateix can not multiply!")
	}
	row, column, middle := len(matA), len(matB[0]), len(matA[0])
	matC := make([][]int, row)
	for i := 0; i < row; i++ {
		matC[i] = make([]int, column)
		for j := 0; j < column; j++ {
			matC[i][j] = 0
			for k := 0; k < middle; k++ {
				matC[i][j] = matC[i][j] + matA[i][k]*matB[k][j]
			}
		}
	}
	return matC
}
