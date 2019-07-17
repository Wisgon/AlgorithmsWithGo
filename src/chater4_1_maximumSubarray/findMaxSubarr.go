package chater4_1_maximumSubarray

import (
	"fmt"
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

//ToDO：设计可控制开几个线程的版本
func FindMaxSubArrayCon(sli []int, leftIdx int, rightIdx int, maxLeftIdxResChan chan int, maxRightIdxResChan chan int, maxSumResChan chan int, recursionNum int, setRecursionNum int) {
	if len(sli) < 2*setRecursionNum {
		//数组太少会导致无限递归
		panic("too small sli!")
	}
	midIdx := int((rightIdx + leftIdx) / 2)
	bts := binaryTreeSum(recursionNum) // 注意这里为什么是计算二叉树的值
	fmt.Printf("bts is %d\n", bts)
	recursionNum += 1
	if bts < setRecursionNum {
		fmt.Printf("going recursion %d\n", recursionNum)
		leftMaxLeftIdxChan, leftMaxRightIdxChan, leftMaxSumChan, rightMaxLeftIdxChan, rightMaxRightIdxChan, rightMaxSumChan :=
			make(chan int), make(chan int), make(chan int), make(chan int), make(chan int), make(chan int)

		go FindMaxSubArrayCon(sli, leftIdx, midIdx, leftMaxLeftIdxChan, leftMaxRightIdxChan, leftMaxSumChan, recursionNum, setRecursionNum)
		go FindMaxSubArrayCon(sli, midIdx+1, rightIdx, rightMaxLeftIdxChan, rightMaxRightIdxChan, rightMaxSumChan, recursionNum, setRecursionNum)

		crossMaxLeftIdx, crossMaxRightIdx, crossMaxSum := findMaxCrossingSubarray(sli, leftIdx, rightIdx, midIdx)
		//crossMaxLeftIdx, crossMaxRightIdx, crossMaxSum := findMaxCrossingSubarrayCon(sli, leftIdx, rightIdx, midIdx)
		leftMaxLeftIdx, leftMaxRightIdx, leftMaxSum, rightMaxLeftIdx, rightMaxRightIdx, rightMaxSum :=
			<-leftMaxLeftIdxChan, <-leftMaxRightIdxChan, <-leftMaxSumChan, <-rightMaxLeftIdxChan, <-rightMaxRightIdxChan, <-rightMaxSumChan

		if leftMaxSum >= rightMaxSum && leftMaxSum >= crossMaxSum {
			maxLeftIdxResChan <- leftMaxLeftIdx
			maxRightIdxResChan <- leftMaxRightIdx
			maxSumResChan <- leftMaxSum
		} else if rightMaxSum >= leftMaxSum && rightMaxSum >= crossMaxSum {
			maxLeftIdxResChan <- rightMaxLeftIdx
			maxRightIdxResChan <- rightMaxRightIdx
			maxSumResChan <- rightMaxSum
		} else {
			maxLeftIdxResChan <- crossMaxLeftIdx
			maxRightIdxResChan <- crossMaxRightIdx
			maxSumResChan <- crossMaxSum
		}
	} else {
		fmt.Println("caculating~~~%d, %d, %d", leftIdx, midIdx, rightIdx)
		leftMaxLeftIdx, leftMaxRightIdx, leftMaxSum := FindMaxSubarray(sli, leftIdx, midIdx)
		rightMaxLeftIdx, rightMaxRightIdx, rightMaxSum := FindMaxSubarray(sli, midIdx+1, rightIdx)
		crossMaxLeftIdx, crossMaxRightIdx, crossMaxSum := findMaxCrossingSubarray(sli, leftIdx, rightIdx, midIdx)
		//crossMaxLeftIdx, crossMaxRightIdx, crossMaxSum := findMaxCrossingSubarrayCon(sli, leftIdx, rightIdx, midIdx)
		if leftMaxSum >= rightMaxSum && leftMaxSum >= crossMaxSum {
			maxLeftIdxResChan <- leftMaxLeftIdx
			maxRightIdxResChan <- leftMaxRightIdx
			maxSumResChan <- leftMaxSum
		} else if rightMaxSum >= leftMaxSum && rightMaxSum >= crossMaxSum {
			maxLeftIdxResChan <- rightMaxLeftIdx
			maxRightIdxResChan <- rightMaxRightIdx
			maxSumResChan <- rightMaxSum
		} else {
			maxLeftIdxResChan <- crossMaxLeftIdx
			maxRightIdxResChan <- crossMaxRightIdx
			maxSumResChan <- crossMaxSum
		}
	}
}

func findMaxCrossingSubarrayCon(sli []int, leftIdx int, rightIdx int, midIdx int) (maxLeftIdx int, maxRightIdx int, maxSum int) {

	//实验证明， 用这个多线程版本会更慢，因为递归创造了太多的goroutine

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

func binaryTreeSum(a int) int {
	sum := 1.0
	for i := 1; i < a; i++ {
		sum += math.Pow(2, float64(i))
	}
	return int(sum)
}
