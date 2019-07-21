package usefulFunc

import (
	"log"
)

func MatrixesCanMultiply(A [][]int, B [][]int) bool {
	if len(A[0]) == len(B) {
		return true
	} else {
		return false
	}
}

func DealWithError(how string, err error) {
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

func CutMatrix(mat [][]int, rowMin int, rowMax int, colMin int, colMax int) [][]int {
	var newMatrix [][]int
	var TheRow []int
	matSliceRow := mat[rowMin:rowMax]
	for i := 0; i < rowMax-rowMin; i++ {
		TheRow = matSliceRow[i]
		newMatrix = append(newMatrix, TheRow[colMin:colMax])
	}
	return newMatrix
}
