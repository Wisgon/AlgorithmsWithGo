package algorithmsTest

import (
	"AlgorithmsWithGo/chater4_1_maximumSubarray"
	"AlgorithmsWithGo/usefulFunc"
	"math"
	"time"

	//"../usefulFunc"
	"fmt"
)

func MaxSubarrayTest() {
	//经测试，开八线程的话，在元素数量达到333333335的情况下，多线程版本大概快3倍多
	sli := usefulFunc.CreateRandSlice(333333335, math.MaxInt32, true)
	//sli := []int{-1, -3, -30, 12, -3, 5, 6, 8, -9, -21, -3, 2, 6, -4, 4, 4, -5, 6, 76,
	//	-5, -45, -66, 5, 6, 5, 4, 4, 4, 5, 6, 3, 3, 2, 3, 3, 4, 5, 6, -66, 5, 6, 5, 4,
	//5, 6, 5, 4, -66, 5, 6, 5, 4, -66, 5, 6, 5, 4, -66, 5, 6, 5, 4, -66, 5, 6, 5, 4} //开八个线程的话，数组元素数量不能太小

	ts := time.Now()
	_, _, _ = chater4_1_maximumSubarray.FindMaxSubarray(sli, 0, len(sli)-1)
	//fmt.Println(sum)
	fmt.Println("single time is:", time.Since(ts))

	//concurrent
	ts = time.Now()
	leftChan, righChan, sumChan := make(chan int, 1), make(chan int, 1), make(chan int, 1)
	chater4_1_maximumSubarray.FindMaxSubArrayCon(sli, 0, len(sli)-1, leftChan, righChan, sumChan, 1, 8)
	//fmt.Println(<-sumChan)
	fmt.Println("concurrent time is:", time.Since(ts))
}
