package main

import (
	"chater4_2_matrixMultiply"
	"fmt"
	"math"
	"time"
	"usefulFunc"
)

func main() {
	//normalAlgorithms
	matA := usefulFunc.CreateMatrix(32, 32, math.MinInt8, math.MaxInt8)
	matB := usefulFunc.CreateMatrix(32, 32, math.MinInt8, math.MaxInt8)
	//matA := [][]int{{1, 2, 3, 0}, {4, 5, 6, 0}, {7,8,9, 0}, {14, 15, 16, 0}}
	//matB := [][]int{{3,4,5,6}, {6,7,8,9}, {9,8,7,6}, {5,4,3,2}}
	ts := time.Now()
	matC := chater4_2_matrixMultiply.SquareMatrixMultiply(matA, matB)
	fmt.Println(matC[3][3], time.Since(ts))
	ts2 := time.Now()
	matAp, matBp := make([][]*int, len(matA)), make([][]*int, len(matB))
	for i := 0; i < len(matA); i++ {
		matAp[i], matBp[i] = make([]*int, len(matA[i])), make([]*int, len(matB[i]))
		for j := 0; j < len(matA[i]); j++ {
			matAp[i][j] = &matA[i][j]
			matBp[i][j] = &matB[i][j]
		}
	}
	matC2 := chater4_2_matrixMultiply.SquareMatrixMultiplyRecursive(matAp, matBp)
	fmt.Println(*matC2[3][3], time.Since(ts2))

}
