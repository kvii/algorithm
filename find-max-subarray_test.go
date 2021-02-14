package main

import "testing"

func TestFindMaxSubarray(t *testing.T) {
	A := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	ans := MaxSubarray{7, 10, 43}

	if info1 := FindMaxSubarray(A, 0, len(A)-1); info1 != ans {
		t.Fatal("info1 failed", info1)
	}

	if info2 := FindMaxSubarray2(A); info2 != ans {
		t.Fatal("info2 failed", info2)
	}
}
