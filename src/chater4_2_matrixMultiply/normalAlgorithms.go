package chater4_2_matrixMultiply

import (
	"usefulFunc"
)

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

//多线程版本
func SquareMatrixMultiplyCon(matA [][]int, matB [][]int, coreNum int) [][]int {
	if !usefulFunc.MatrixesCanMultiply(matA, matB) {
		panic("The two mateix can not multiply!")
	}
	var a = len(matA) / coreNum
	var perNum = int(a)
	matArray := make([][][]int, 0, coreNum)
	for i := 0; i < coreNum-1; i++ {
		matArray = append(matArray, matA[i*perNum:(i+1)*perNum])
	}
	matArray = append(matArray, matA[(coreNum-1)*perNum:])

	resArray := make([]chan [][]int, 0, coreNum)
	for i := 0; i < coreNum; i++ {
		resArray = append(resArray, make(chan [][]int, 1))
	}
	for i := 0; i < coreNum; i++ {
		go SquareMatrixMultiplyChan(matArray[i], matB, resArray[i])
	}

	matC := make([][]int, 0, coreNum)
	for i := 0; i < coreNum; i++ {
		matC = append(matC, <-resArray[i]...)
	}
	return matC
}

func SquareMatrixMultiplyChan(matA [][]int, matB [][]int, chanC chan [][]int) {
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
	chanC <- matC
}
