package usefulFunc

import (
	"math/rand"
	"time"
)

func CreateRandSlicePositive(amount int) []int {
	//这个函数创造amount个全是正整数的随机数
	var s = make([]int, amount)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < amount; i++ {
		s[i] = rand.Int()
	}
	return s
}

func CreateRandSlice(amount int, biggestNum int, ifNegitive bool) []int {
	/*这个函数创造amount个整数，最大不超过biggestNUm，如果ifNegitive是true，那么就会有负数生成，
	这时候正数的最大数不会超过biggestNum/2，所以一般如果想生成int64类型的正负数都有的sllice，那么biggestNum
	必须传入math.MaxInt32*/
	var s = make([]int, amount)
	rand.Seed(time.Now().UnixNano())
	if ifNegitive {
		for i := 0; i < amount; i++ {
			s[i] = rand.Intn(biggestNum*2) - biggestNum
		}
	} else {
		for i := 0; i < amount; i++ {
			s[i] = rand.Intn(biggestNum)
		}
	}
	return s
}

func CreateMatrix(row int, column int, minNum int, maxNum int) [][]int {
	//math.MaxInt8可获得8位整数的最大值
	if minNum >= maxNum {
		panic("minNum must little than maxNum!")
	}
	var newMatrix [][]int
	newMatrix = make([][]int, row)
	rand.Seed(time.Now().Unix())
	for i := 0; i < row; i++ {
		newMatrix[i] = make([]int, column)
		for j := 0; j < column; j++ {
			newMatrix[i][j] = rand.Intn(maxNum-minNum) + minNum
		}
	}
	return newMatrix
}
