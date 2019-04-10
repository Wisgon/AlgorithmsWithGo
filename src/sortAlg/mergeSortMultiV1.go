package sortAlg

import (
	"sync"
)

func MergeSortMultiOne(sli []int, chunkNum int) {
	//这个归并排序是用通道实现的
	var container [][]int
	var wg sync.WaitGroup

	LengthOfSli := len(sli)
	if LengthOfSli <= chunkNum {
		panic("Too many chunks")
	}

	Aer := int(LengthOfSli / chunkNum) // Aer = AmountEachRoutine
	for i := 0; i < chunkNum-1; i++ {
		container = append(container, sli[i*Aer:(i+1)*Aer])
		wg.Add(1)
	}
	container = append(container, sli[(chunkNum-1)*Aer:])
	wg.Add(1)
	//ts := time.Now()
	for _, v := range container {
		go func(v []int) {
			MergeSort(v, 0, len(v)-1)
			wg.Done()
		}(v)
	}

	wg.Wait()
	//fmt.Println("sort done", time.Since(ts))

	elementNum := 0
	for _, v := range container {
		elementNum += len(v)
	}
	result := make(chan int, elementNum+10)
	go mergeMulti(container, result)

	for i, _ := range sli {
		sli[i] = <-result
	}

}

func mergeMulti(container [][]int, result chan int) {

	defer close(result) // 如果在41行result的声明处close掉，就会产生死锁
	if len(container) == 1 {
		for _, v := range container[0] {
			result <- v
		}
	} else {
		splitContainer := int(len(container) / 2)
		e1, e2 := 0, 0
		for _, v := range container[0:splitContainer] {
			e1 += len(v)
		}
		for _, v := range container[splitContainer:] {
			e2 += len(v)
		}
		res1, res2 := make(chan int, e1+10), make(chan int, e2+10)

		go mergeMulti(container[0:splitContainer], res1)
		go mergeMulti(container[splitContainer:], res2)

		v1, ok1 := <-res1
		v2, ok2 := <-res2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				result <- v1
				v1, ok1 = <-res1
			} else {
				result <- v2
				v2, ok2 = <-res2
			}
		}
	}
}
