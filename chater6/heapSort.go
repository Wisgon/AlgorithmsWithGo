package chater6

func Parent(i int) int {
	/*
		params:
			i: the index of an element in the array
		returns:
			the index of parent of the element in the tree of heap
	*/
	return i >> 1
}

func Left(i int) int {
	/*
		params:
			i: the index of an element in the array
		returns:
			the index of left children of the element in the tree of heap
	*/
	return i << 1
}

func Right(i int) int {
	return i<<1 + 1
}

func MaxHeapIfy(arr *[]int, i int) {
	// implement of chater 6.2 MAX-HEAPIFY
	/*
		params:
			arr: pointer of arr that need to sort
			i: index of node plus one
		returns:
			no returns
	*/
	var largestIndex int
	l := Left(i)
	r := Right(i)
	// index must minux 1 because array is start from 0
	if l <= len(*arr) && (*arr)[l-1] > (*arr)[i-1] {
		largestIndex = l
	} else {
		largestIndex = i
	}
	if r <= len(*arr) && (*arr)[r-1] > (*arr)[largestIndex-1] {
		largestIndex = r
	}
	if largestIndex != i {
		// exchange A[i] with A[largest]
		(*arr)[largestIndex-1], (*arr)[i-1] = (*arr)[i-1], (*arr)[largestIndex-1]
		MaxHeapIfy(arr, largestIndex)
	}
}

func BuildMaxHeap(arr *[]int) {
	// implement of chater 6.3 BUILD-MAX-HEAP
	for i := len(*arr)>>1 - 1; i >= 0; i-- {
		// i must plus 1 because MaxHeapIfy think that arr is start on 1
		MaxHeapIfy(arr, i+1)
	}
}

func HeapSort(arr *[]int, sorted *[]int) {
	// implement of chater 6.4 HEAPSORT
	BuildMaxHeap(arr)
	for i := len(*arr) - 1; i > 0; i-- {
		// fmt.Println("arr:", *arr, "sorted", *sorted)
		(*arr)[0], (*arr)[i] = (*arr)[i], (*arr)[0]
		(*sorted) = append((*sorted), (*arr)[i])
		*arr = (*arr)[:i]
		MaxHeapIfy(arr, 1)
	}
	(*sorted) = append((*sorted), (*arr)[0]) // must add the last one
}
