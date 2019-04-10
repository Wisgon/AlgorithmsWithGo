package main

import (
	"../maximumSubarray"
	"math"
	"time"
	"usefulFunc"

	//"../usefulFunc"
	"fmt"
)

func main() {
	sli := usefulFunc.CreateRandSlice(1333, math.MaxInt32, true)
	//sli := []int{-1, -3, -30, 12, -3, 5, 6, 8, -9, -21, -3, 2}
	ts := time.Now()
	_, _, _ = maximumSubarray.FindMaxSubarray(sli, 0, len(sli)-1)
	//left, right, sum := maximumSubarray.FindMaxSubarray(sli, 0, len(sli)-1)
	fmt.Println("time is:", time.Since(ts))
	//fmt.Println(sli[left:right+1], sum)
}
