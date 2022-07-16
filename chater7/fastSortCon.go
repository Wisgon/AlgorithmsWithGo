package chater7

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func RunFastSortCon() {
	A := []int{3, 7, 4, 3, 9, 2, 1, 8}
	workerNumb := 4
	wg.Add(1)
	go QuickSortCon(A, 0, 7, 1, workerNumb)
	wg.Wait()
	fmt.Println("A:", A)
}

func QuickSortCon(A []int, p int, r int, goRoutineCount int, workerNumb int) {
	if p < r {
		q := PartitionCon(A, p, r)

		if goRoutineCount >= workerNumb { // 开workerNumb个线程
			QuickSortCon(A, p, q-1, goRoutineCount+2, workerNumb)
			QuickSortCon(A, q+1, r, goRoutineCount+2, workerNumb)
			fmt.Println("goroucou", goRoutineCount)
		} else {
			goRoutineCount *= 2
			fmt.Println("starting a con:", goRoutineCount)
			wg.Add(1)
			go QuickSortCon(A, p, q-1, goRoutineCount, workerNumb)
			wg.Add(1)
			go QuickSortCon(A, q+1, r, goRoutineCount, workerNumb)
		}
	}
	if goRoutineCount == 1 || goRoutineCount <= workerNumb {
		// 当goRoutineCount为1为一开始的线程，为小于或等于workerNumb为开了goroutine的线程
		wg.Done()
	}
}

func PartitionCon(A []int, p int, r int) int {
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
