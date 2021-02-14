package main

import (
	"math"
)

// MaxSubarray is function find_max_subarray's answer
type MaxSubarray struct {
	left  int // 极左下标
	right int // 极右下标
	sum   int // Δ合计
}

// FindMaxSubarray find max subarray in A[0..len(A)-1]
func FindMaxSubarray(A []int, low, high int) (ans MaxSubarray) {
	if low == high {
		return MaxSubarray{low, high, A[low]}
	}

	mid := (low + high) / 2
	var (
		info1 = FindMaxSubarray(A, low, mid)
		info2 = FindMaxSubarray(A, mid+1, high)
		info3 = findMaxCrossingSubarray(A, low, mid, high)
	)
	if info1.sum >= info2.sum && info1.sum >= info3.sum {
		return info1
	}
	if info2.sum >= info1.sum && info2.sum >= info3.sum {
		return info2
	}
	return info3
}

func findMaxCrossingSubarray(A []int, low, mid, high int) (ans MaxSubarray) {

	var (
		leftSum = int(math.Inf(-1))
		sum1    = 0
	)
	for i := mid; i >= low; i-- {
		sum1 += A[i]
		if sum1 > leftSum {
			leftSum = sum1
			ans.left = i
		}
	}

	var (
		rightSum = int(math.Inf(-1))
		sum2     = 0
	)
	for j := mid + 1; j <= high; j++ {
		sum2 += A[j]
		if sum2 > rightSum {
			rightSum = sum2
			ans.right = j
		}
	}

	ans.sum = leftSum + rightSum
	return
}

// FindMaxSubarray2 answer for 4.1-1
func FindMaxSubarray2(A []int) (ans MaxSubarray) {
	for left := range A {
		var sum int
		for right := left; right < len(A); right++ {
			sum += A[right]
			if sum > ans.sum {
				ans = MaxSubarray{left, right, sum}
			}
		}
	}
	return
}
