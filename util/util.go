package util

import (
	"math/rand"
)

const maxDelta = 1000

// MockSliceInt generate an int slice
func MockSliceInt(n int) []int {

	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		ans = append(ans, randInt())
	}
	return ans
}

// [-delta, +delta]
func randInt() int { return rand.Intn(1+2*maxDelta) - maxDelta }
