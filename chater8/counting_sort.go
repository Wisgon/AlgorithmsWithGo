package chater8

import "AlgorithmsWithGo/usefulFunc"

func CountingSort(targetArray []int) []int {
	_, maxValue := usefulFunc.FindMaxElement(targetArray)
	tmpArray := make([]int, maxValue+1)
	for _, v := range targetArray {
		tmpArray[v] = tmpArray[v] + 1
	}
	for i := 1; i < len(tmpArray); i++ {
		tmpArray[i] = tmpArray[i] + tmpArray[i-1]
	}
	resultArray := make([]int, len(targetArray))
	for i := 0; i < len(targetArray); i++ {
		resultArray[tmpArray[targetArray[i]]-1] = targetArray[i]
		tmpArray[targetArray[i]] = tmpArray[targetArray[i]] - 1
	}
	return resultArray
}
