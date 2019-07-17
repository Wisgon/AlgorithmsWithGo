package main

import (
	"chater4_1_maximumSubarray"
	"time"
	//"../usefulFunc"
	"fmt"
)

func main() {
	//sli := usefulFunc.CreateRandSlice(1333, math.MaxInt32, true)
	sli := []int{-1, -3, -30, 12, -3, 5, 6, 8, -9, -21, -3, 2, 6, 4, 4, 4, 5, 6, 76, -5, -45, -66, 5, 6, 5, 4, 4, 4, 5, 6, 3, 3, 2, 3, 3, 4, 5, 6} //开八个线程的话，数组元素数量不能太小
	ts := time.Now()
	//left, right, sum := chater4_1_maximumSubarray.FindMaxSubarray(sli, 0, len(sli)-1)
	leftChan, righChan, sumChan := make(chan int), make(chan int), make(chan int)
	chater4_1_maximumSubarray.FindMaxSubArrayCon(sli, 0, len(sli)-1, leftChan, righChan, sumChan, 1, 8)
	fmt.Println("time is:", time.Since(ts))
	//fmt.Println(sli[left:right+1], sum)
	fmt.Println(sli[<-leftChan:<-righChan+1], <-sumChan)
}
