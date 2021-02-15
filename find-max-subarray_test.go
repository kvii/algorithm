package main

import (
	"algorithm/util"
	"math"
	"testing"
	"time"
)

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

func Test413(t *testing.T) {
	for n := 1; n < math.MaxInt64; n++ {
		A := util.MockSliceInt(n)

		t1 := time.Now()
		info1 := FindMaxSubarray(A, 0, len(A)-1)
		d1 := time.Since(t1)

		t2 := time.Now()
		info2 := FindMaxSubarray2(A)
		d2 := time.Since(t2)

		if info1 != info2 {
			t.Error("err")
			continue
		}
		t.Log(n, d1, d2)
		if d1 < d2 {
			t.SkipNow()
		}
	}
}
