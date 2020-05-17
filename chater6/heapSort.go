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

func MaxHeapIfy(arr *[]int) {
	// implement of chater 6.2 MAX-HEAPIFY
	
}
