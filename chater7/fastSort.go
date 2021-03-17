package main

import "fmt"

func main() {
	A := []int{3, 7, 4, 3, 9, 2, 1, 8}
	QuickSort(A, 0, 7)
	fmt.Println("A:", A)
}

func QuickSort(A []int, p int, r int) {
	if p < r {
		q := Partition(A, p, r)
		QuickSort(A, p, q-1)
		QuickSort(A, q+1, r)
	}
}

func Partition(A []int, p int, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if A[j] <= x {
			i += 1
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}
