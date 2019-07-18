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
	matA := usefulFunc.CreateMatrix(100, 100, math.MinInt8, math.MaxInt8)
	matB := usefulFunc.CreateMatrix(100, 100, math.MinInt8, math.MaxInt8)
	ts := time.Now()
	matC := chater4_2_matrixMultiply.SquareMatrixMultiply(matA, matB)
	fmt.Println(matC[33][55], time.Since(ts))

}
