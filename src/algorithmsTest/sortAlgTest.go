package main

import (
	"fmt"
	"sortAlg"
	"sync"
	"time"
	"usefulFunc"
)

func main() {
	data := usefulFunc.CreateRandSlicePositive(49801245)
	copyData := make([]int, len(data)) //用来存放生成的数组，相同的数组才有比较的意义
	ts := time.Now()

	//copy(copyData, data)
	//ts = time.Now()
	//sortAlg.BaboSort(copyData) //同一个机器十万个排序花了15秒
	//fmt.Println("time one is ", time.Since(ts))
	//usefulFunc.ShowSlice(copyData)

	//copy(copyData, data)
	//ts = time.Now()
	//sortAlg.InsertionSort(copyData) //同一个机器十万个排序花了2.1秒
	//fmt.Println("time two is", time.Since(ts))
	//usefulFunc.ShowSlice(copyData)

	copy(copyData, data)
	ts = time.Now()
	//单线程归并排序
	sortAlg.MergeSort(copyData, 0, len(copyData)-1) //同一个机器十万个排序花了四十毫秒
	fmt.Println("time three is", time.Since(ts))
	//usefulFunc.ShowSlice(copyData)

	//copy(copyData, data)
	////usefulFunc.ShowSlice(copyData)
	//ts = time.Now()
	////这个是用通道实现的归并排序，估计通道通信消耗的时间太多，所以最慢
	//sortAlg.MergeSortMultiOne(copyData, 8) //同一个机器十万个排序花了八十四毫秒
	//fmt.Println("time four is", time.Since(ts))
	//usefulFunc.ShowSlice(copyData)

	copy(copyData, data)
	//usefulFunc.ShowSlice(copyData)
	ts = time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	//控制住多个线程的归并排序，chunkNum设置太多会发生内存爆炸
	go sortAlg.MergeSortMultitwo(copyData, 0, len(copyData)-1, &wg, 8) //同一个机器十万个排序花了十三毫秒
	wg.Wait()
	fmt.Println("time five is", time.Since(ts))
	//usefulFunc.ShowSlice(copyData)

}
