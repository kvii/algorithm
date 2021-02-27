package main

// SquareMatrixMultiply 矩阵乘法
/*
	QUARE_MATRIX_MULTIPLY(A, B) // n = A.rows
		C = make_matrix(n, n)
		for i := 1 to n
			for j := 1 to n
				C[i][j] = 0
				for k := 1 to n
					C[i][j] = C[i][j] + A[i][k] * B[k][j]
		return C
*/
func SquareMatrixMultiply(A, B [][]int) (C [][]int) {
	C = make([][]int, len(A))
	for i := range A {
		C[i] = make([]int, len(A))
		for j := range A {
			for k := range A {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return
}
