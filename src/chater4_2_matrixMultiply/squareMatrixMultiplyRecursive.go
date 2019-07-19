package chater4_2_matrixMultiply

import (
	uf "usefulFunc"
)

func SquareMatrixMultiplyRecursive(matA [][]int, matB [][]int) [][]int {
	n1, n2, n3, n4 := len(matA), len(matA[0]), len(matB), len(matB[0])
	if !(uf.IsPowOfTwo(n1) || uf.IsPowOfTwo(n2) || uf.IsPowOfTwo(n3) || uf.IsPowOfTwo(n4)) || n1 != n2 || n2 != n3 || n3 != n4 {
		panic("Matrix must be n x n and n must be pow of 2")
	}
	matC := make([][]int, n1)
	if n1 == 1 {
		matC[0][0] = matA[0][0] * matB[0][0]
	} else {

	}

}

func MatrixPlus(matA [][]int, matB [][]int) [][]int {
	matC := make([][]int, len(matA))
	for i := 0; i < len(matA); i++ {
		matC[i] = make([]int, len(matA[i]))
		for j := 0; j < len(matA[i]); j++ {
			matC[i][j] = matA[i][j] + matB[i][j]
		}
	}
	return matC
}
