// sortAlgorithms集合了各章节的排序算法
package algorithmsTest

import (
	"AlgorithmsWithGo/sortAlgorithms"
	"fmt"
	"sync"
	"time"
	"usefulFunc"
)

var data = usefulFunc.CreateRandSlicePositive(100000)

func BaboSortTest() {
	copyData := make([]int, len(data)) //用来存放生成的数组，相同的数组才有比较的意义
	ts := time.Now()
	copy(copyData, data)
	ts = time.Now()
	sortAlgorithms.BaboSort(copyData) //同一个机器十万个排序花了15秒
	fmt.Println("baboSort time is ", time.Since(ts))
	usefulFunc.ShowSlice(copyData)
}

func InsertionSortTest() {
	copyData := make([]int, len(data)) //用来存放生成的数组，相同的数组才有比较的意义
	ts := time.Now()
	copy(copyData, data)
	ts = time.Now()
	sortAlgorithms.InsertionSort(copyData) //同一个机器十万个排序花了2.1秒
	fmt.Println("time two is", time.Since(ts))
	usefulFunc.ShowSlice(copyData)
}

func MergeSortTest() {
	copyData := make([]int, len(data)) //用来存放生成的数组，相同的数组才有比较的意义
	ts := time.Now()
	copy(copyData, data)
	ts = time.Now()
	//单线程归并排序
	sortAlgorithms.MergeSort(copyData, 0, len(copyData)-1) //同一个机器十万个排序花了四十毫秒
	fmt.Println("time three is", time.Since(ts))
	//usefulFunc.ShowSlice(copyData)
}

func MergeSortMultiOneTest() {
	copyData := make([]int, len(data)) //用来存放生成的数组，相同的数组才有比较的意义
	ts := time.Now()
	copy(copyData, data)
	//usefulFunc.ShowSlice(copyData)
	ts = time.Now()
	//这个是用通道实现的归并排序，估计通道通信消耗的时间太多，所以最慢
	sortAlgorithms.MergeSortMultiOne(copyData, 8) //同一个机器十万个排序花了八十四毫秒
	fmt.Println("time four is", time.Since(ts))
	usefulFunc.ShowSlice(copyData)
}

func MergeSortMultitwoTest() {
	copyData := make([]int, len(data)) //用来存放生成的数组，相同的数组才有比较的意义
	ts := time.Now()
	copy(copyData, data)
	//usefulFunc.ShowSlice(copyData)
	ts = time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	//控制住多个线程的归并排序，chunkNum设置太多会发生内存爆炸
	go sortAlgorithms.MergeSortMultitwo(copyData, 0, len(copyData)-1, &wg, 4) //同一个机器十万个排序花了十三毫秒
	wg.Wait()
	fmt.Println("time five is", time.Since(ts))
	//usefulFunc.ShowSlice(copyData)
}
