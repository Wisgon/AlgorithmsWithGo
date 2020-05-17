package usefulFunc

import (
	"fmt"
	"log"
)

func ShowSlice(sli []int) {
	for _, v := range sli {
		fmt.Println(v)
	}
}

func MatrixesCanMultiply(A [][]int, B [][]int) bool {
	// 判断两个矩阵是否可以相乘
	if len(A[0]) == len(B) {
		return true
	} else {
		return false
	}
}

func DealWithError(how string, err error) {
	/*
	usage: 处理error的不同方式的封装
	param:
		how:如何来处理错误，可以使字符串"logln"，表示用log.Fatalln(err)
		err:错误对象
	return: 无
	*/
	if err != nil {
		switch how {
		case "logln":
			log.Fatalln(err)
		case "panic":
			panic(err)
		default:
			panic("Unknow method to deal with error!")
		}
	}
}

func IsPowOfTwo(a int) bool {
	//二的幂全是0，和比自己小1的数按位与得0
	return a&(a-1) == 0
}

func CutMatrix(mat [][]*int, rowMin int, rowMax int, colMin int, colMax int) [][]*int {
	//这个函数是将矩阵mat的rowMin行到rowMax行，colMin列到colMax列的元素抽出来
	newMatrix := make([][]*int, rowMax-rowMin)
	matSliceRow := mat[rowMin:rowMax]
	for i := 0; i < rowMax-rowMin; i++ {
		newMatrix[i] = make([]*int, colMax-colMin)
		theRow := matSliceRow[i][colMin:colMax]
		for j := 0; j < colMax-colMin; j++ {
			newMatrix[i][j] = theRow[j]
		}
	}
	return newMatrix
}
