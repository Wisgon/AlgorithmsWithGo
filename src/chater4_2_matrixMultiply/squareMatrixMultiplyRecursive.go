package chater4_2_matrixMultiply

import (
	uf "usefulFunc"
)

func SquareMatrixMultiplyRecursive(matA [][]int, matB [][]int) [][]int {
	n1, n2, n3, n4 := len(matA), len(matA[0]), len(matB), len(matB[0])
	if !(uf.IsPowOfTwo(n1) || uf.IsPowOfTwo(n2) || uf.IsPowOfTwo(n3) || uf.IsPowOfTwo(n4)) || n1 != n2 || n2 != n3 || n3 != n4 {
		panic("Matrix must both be n x n matrix and n must be pow of 2")
	}
	if n1 == 1 {
		return [][]int{{matA[0][0] * matB[0][0]}}
	} else {
		//C11 := MatrixPlus(
		//	SquareMatrixMultiplyRecursive(matA[:n1/2][:n1/2], matB[:n1/2][:n1/2]),
		//	SquareMatrixMultiplyRecursive(matA[:n1/2][n1/2:], matB[n1/2:][:n1/2]),
		//)
		//C12 := MatrixPlus(
		//	SquareMatrixMultiplyRecursive(matA[:n1/2][:n1/2], matB[:n1/2][n1/2:]),
		//	SquareMatrixMultiplyRecursive(matA[:n1/2][n1/2:], matB[n1/2:][n1/2:]),
		//)
		//C13 := MatrixPlus(
		//	SquareMatrixMultiplyRecursive(matA[n1/2:][:n1/2], matB[:n1/2][:n1/2]),
		//	SquareMatrixMultiplyRecursive(matA[n1/2:][n1/2:], matB[n1/2:][:n1/2]),
		//)
		//C14 := MatrixPlus(
		//	SquareMatrixMultiplyRecursive(matA[n1/2:][:n1/2], matB[:n1/2][n1/2:]),
		//	SquareMatrixMultiplyRecursive(matA[n1/2:][n1/2:], matB[n1/2:][n1/2:]),
		//)
		C11 := MatrixPlus(
			SquareMatrixMultiplyRecursive(
				uf.CutMatrix(matA, 0, n1/2, 0, n1/2),
				uf.CutMatrix(matB, 0, n1/2, 0, n1/2),
			),
			SquareMatrixMultiplyRecursive(
				uf.CutMatrix(matA, 0, n1/2, n1/2, n1),
				uf.CutMatrix(matB, n1/2, n1, 0, n1/2),
			),
		)
		C12 := MatrixPlus(
			SquareMatrixMultiplyRecursive(
				uf.CutMatrix(matA, 0, n1/2, 0, n1/2),
				uf.CutMatrix(matB, 0, n1/2, n1/2, n1),
			),
			SquareMatrixMultiplyRecursive(
				uf.CutMatrix(matA, 0, n1/2, n1/2, n1),
				uf.CutMatrix(matB, n1/2, n1, n1/2, n1),
			),
		)
		C13 := MatrixPlus(
			SquareMatrixMultiplyRecursive(
				uf.CutMatrix(matA, n1/2, n1, 0, n1/2),
				uf.CutMatrix(matB, 0, n1/2, 0, n1/2),
			),
			SquareMatrixMultiplyRecursive(
				uf.CutMatrix(matA, n1/2, n1, n1/2, n1),
				uf.CutMatrix(matB, n1/2, n1, 0, n1/2),
			),
		)
		C14 := MatrixPlus(
			SquareMatrixMultiplyRecursive(
				uf.CutMatrix(matA, n1/2, n1, 0, n1/2),
				uf.CutMatrix(matB, 0, n1/2, n1/2, n1),
			),
			SquareMatrixMultiplyRecursive(
				uf.CutMatrix(matA, n1/2, n1, n1/2, n1),
				uf.CutMatrix(matB, n1/2, n1, n1/2, n1),
			),
		)

		return CombineMatrix(C11, C12, C13, C14)
	}

}

func MatrixPlus(matA [][]int, matB [][]int) [][]int {
	//两个矩阵对应元素相加的操作
	matC := make([][]int, len(matA))
	for i := 0; i < len(matA); i++ {
		matC[i] = make([]int, len(matA[i]))
		for j := 0; j < len(matA[i]); j++ {
			matC[i][j] = matA[i][j] + matB[i][j]
		}
	}
	return matC
}

func CombineMatrix(matA [][]int, matB [][]int, matC [][]int, matD [][]int) [][]int {
	//将4个nxn矩阵组成一个大的2nx2n的矩阵
	n := len(matA)
	var mat [][]int
	for i := 0; i < n; i++ {
		mat = append(mat, append(matA[i], matB[i]...))
	}
	for j := 0; j < n; j++ {
		mat = append(mat, append(matC[j], matD[j]...))
	}
	return mat
}
