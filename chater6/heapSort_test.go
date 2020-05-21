package chater6

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{23, 66, 54, 6, 78, 334, 98, 21, 0, 21, 6, 3, 5, 345, 3445, 88, 88, 46, 67, 77, 5, 655, 6}
	sorted := make([]int, 0, len(arr))
	HeapSort(&arr, &sorted)
	fmt.Println(sorted)
}

func TestMaxHeapIfy(t *testing.T) {
	arr := []int{3, 6, 5, 6, 4}
	MaxHeapIfy(&arr, 1)
	fmt.Println("arr:", arr)
}
