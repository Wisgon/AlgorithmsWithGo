package chater4_1_maximumSubarray

import (
	"math"
)

func FindMaxSubarray(sli []int, leftIdx int, rightIdx int) (maxLeftIdx int, maxRightIdx int, maxSum int) {
	//这里一开始的rightIdx必须是len(sli)-1
	if leftIdx == rightIdx {
		return leftIdx, rightIdx, sli[leftIdx]
	} else {
		midIdx := int((rightIdx + leftIdx) / 2)
		leftMaxLeftIdx, leftMaxRightIdx, leftMaxSum := FindMaxSubarray(sli, leftIdx, midIdx)
		rightMaxLeftIdx, rightMaxRightIdx, rightMaxSum := FindMaxSubarray(sli, midIdx+1, rightIdx)
		crossMaxLeftIdx, crossMaxRightIdx, crossMaxSum := findMaxCrossingSubarray(sli, leftIdx, rightIdx, midIdx)
		//crossMaxLeftIdx, crossMaxRightIdx, crossMaxSum := findMaxCrossingSubarrayCon(sli, leftIdx, rightIdx, midIdx)
		if leftMaxSum >= rightMaxSum && leftMaxSum >= crossMaxSum {
			return leftMaxLeftIdx, leftMaxRightIdx, leftMaxSum
		} else if rightMaxSum >= leftMaxSum && rightMaxSum >= crossMaxSum {
			return rightMaxLeftIdx, rightMaxRightIdx, rightMaxSum
		} else {
			return crossMaxLeftIdx, crossMaxRightIdx, crossMaxSum
		}
	}
}

func findMaxCrossingSubarrayCon(sli []int, leftIdx int, rightIdx int, midIdx int) (maxLeftIdx int, maxRightIdx int, maxSum int) {

	//实验证明， 用这个多线程版本会更慢，因为递归创造了太多的goroutine
	//ToDO：设计最多只开8个线程的版本

	maxLeftIdxChan, maxRightIdxChan, maxLeftSumChan, maxRightSumChan := make(chan int), make(chan int), make(chan int), make(chan int)

	go func(sli []int, leftIdx int, midIdx int) {
		leftSum := -math.MaxInt64
		sum := 0
		var left int
		for i := midIdx; i >= leftIdx; i-- {
			sum += sli[i]
			if sum >= leftSum {
				leftSum = sum
				left = i
			}
		}
		maxLeftIdxChan <- left
		maxLeftSumChan <- leftSum
	}(sli, leftIdx, midIdx)

	go func(sli []int, rightIdx int, midIdx int) {
		rightSum := -math.MaxInt64
		sum := 0
		var right int
		for i := midIdx + 1; i <= rightIdx; i++ {
			sum += sli[i]
			if sum >= rightSum {
				rightSum = sum
				right = i
			}
		}
		maxRightIdxChan <- right
		maxRightSumChan <- rightSum
	}(sli, rightIdx, midIdx)

	maxLeftIdx = <-maxLeftIdxChan
	maxRightIdx = <-maxRightIdxChan
	maxSum = <-maxLeftSumChan + <-maxRightSumChan
	return maxLeftIdx, maxRightIdx, maxSum
}

func findMaxCrossingSubarray(sli []int, leftIdx int, rightIdx int, midIdx int) (maxLeftIdx int, maxRightIdx int, maxSum int) {
	leftSum := -math.MaxInt64
	rightSum := -math.MaxInt64

	sum := 0
	for i := midIdx; i >= leftIdx; i-- {
		sum += sli[i]
		if sum >= leftSum {
			leftSum = sum
			maxLeftIdx = i
		}
	}

	sum = 0
	for i := midIdx + 1; i <= rightIdx; i++ {
		sum += sli[i]
		if sum >= rightSum {
			rightSum = sum
			maxRightIdx = i
		}
	}

	return maxLeftIdx, maxRightIdx, leftSum + rightSum

}
