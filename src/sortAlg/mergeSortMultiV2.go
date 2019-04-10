package sortAlg

import (
	"math"
	"sync"
)

func MergeSortMultitwo(sli []int, left int, right int, wg *sync.WaitGroup, chunkNum int) {
	// first right must be len(sli)-1
	if left < right {
		middle := int((right + left) / 2)
		if chunkNum > 0 {
			var wg2 sync.WaitGroup
			wg2.Add(2)
			chunkNum -= 2
			go MergeSortMultitwo(sli, left, middle, &wg2, chunkNum)
			go MergeSortMultitwo(sli, middle+1, right, &wg2, chunkNum)
			wg2.Wait()
			mergeMulti2(sli, left, right, middle)
		} else {
			var wg2 sync.WaitGroup
			wg2.Add(2)
			MergeSortMultitwo(sli, left, middle, &wg2, chunkNum)
			MergeSortMultitwo(sli, middle+1, right, &wg2, chunkNum)
			mergeMulti2(sli, left, right, middle)
		}
	}
	wg.Done()
}

func mergeMulti2(sli []int, left int, right int, middle int) {

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
