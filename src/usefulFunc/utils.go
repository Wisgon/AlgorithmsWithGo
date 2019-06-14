package usefulFunc

func MatrixesCanMultiply(A [][]int, B [][]int) bool {
	if len(A[0]) == len(B) {
		return true
	} else {
		return false
	}
}
