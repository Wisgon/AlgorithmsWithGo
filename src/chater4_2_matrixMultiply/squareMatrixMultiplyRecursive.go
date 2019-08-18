package chater4_2_matrixMultiply

import (
	uf "gitee.com/wisgon/AlgorithmsWithGo/src/usefulFunc"
)

func SquareMatrixMultiplyRecursive(matA [][]*int, matB [][]*int) [][]*int {
	n1, n2, n3, n4 := len(matA), len(matA[0]), len(matB), len(matB[0])
	//fmt.Printf("matA is %v\n", matA)
	//fmt.Printf("matB is %v\n", matB)
	if !(uf.IsPowOfTwo(n1) || uf.IsPowOfTwo(n2) || uf.IsPowOfTwo(n3) || uf.IsPowOfTwo(n4)) || n1 != n2 || n2 != n3 || n3 != n4 {
		panic("Matrix must both be n x n matrix and n must be pow of 2")
	}
	if n1 == 1 {
		res := *matA[0][0] * *matB[0][0]
		matC := [][]*int{{&res}}
		return matC
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

//下面的两个函数都可以做成append版本，即不用make创建数组，
// 而是用空数组，然后append元素进去，但是append会消耗更多资源，
//因为要重新分配内存一次，所以append的方法更慢
//而用make创建非指针数组之后再赋值，会发生元素复制，虽然比append快20%左右
//将传入数组类型换成指针，也没什么效果

func MatrixPlus(matA [][]*int, matB [][]*int) [][]int {
	//两个矩阵对应元素相加的操作
	matC := make([][]int, len(matA))
	for i := 0; i < len(matA); i++ {
		matC[i] = make([]int, len(matA[i]))
		for j := 0; j < len(matA[i]); j++ {
			matC[i][j] = *matA[i][j] + *matB[i][j]
		}
	}
	return matC
}

func CombineMatrix(matA [][]int, matB [][]int, matC [][]int, matD [][]int) [][]*int {
	//将4个nxn矩阵组成一个大的2nx2n的矩阵
	n := len(matA)
	mat := make([][]*int, 2*n)
	for i := 0; i < n; i++ {
		mat[i] = make([]*int, 2*n)
		mat[i+n] = make([]*int, 2*n)
		for j := 0; j < n; j++ {
			mat[i][j] = &matA[i][j]
			mat[i][j+n] = &matB[i][j]
			mat[i+n][j] = &matC[i][j]
			mat[i+n][j+n] = &matD[i][j]
		}
	}
	return mat
}
