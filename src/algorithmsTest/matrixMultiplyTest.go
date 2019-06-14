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
	matA := usefulFunc.CreateMatrix(1000, 400, math.MinInt8, math.MaxInt8)
	matB := usefulFunc.CreateMatrix(400, 1000, math.MinInt8, math.MaxInt8)
	ts := time.Now()
	matC := chater4_2_matrixMultiply.SquareMatrixMultiply(matA, matB)
	fmt.Println(len(matC), time.Since(ts))

}
