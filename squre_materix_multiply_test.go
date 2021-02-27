package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSqureMaterixMultiply(t *testing.T) {
	A := [][]int{
		{1, 2},
		{3, 4},
	}
	B := [][]int{
		{1, 2},
		{3, 4},
	}
	C := [][]int{
		{7, 10},
		{15, 22},
	}

	C2 := SquareMatrixMultiply(A, B)
	for i, c := range C2 {
		for j, v := range c {
			require.Equal(t, C[i][j], v)
		}
	}
}
