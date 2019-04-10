package sortAlg

import "math"

func MergeSort(sli []int, left int, right int) {
	// first right must be len(sli)-1
	if left < right {
		middle := int((right + left) / 2)
		MergeSort(sli, left, middle)
		MergeSort(sli, middle+1, right)
		merge(sli, left, right, middle)
	}
}

func merge(sli []int, left int, right int, middle int) {

	leftArr, rightArr := make([]int, middle+1-left+1), make([]int, right-middle+1)
	copy(leftArr, sli[left:middle+1])
	copy(rightArr, sli[middle+1:right+1])
	leftArr[middle+1-left], rightArr[right-middle] = math.MaxInt64, math.MaxInt64

	i, j := 0, 0
	for k := left; k < right+1; k++ {
		if leftArr[i] < rightArr[j] {
			sli[k] = leftArr[i]
			i++
		} else {
			sli[k] = rightArr[j]
			j++
		}
	}

}
