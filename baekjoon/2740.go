package main

import "fmt"

func matrixMultiplication(A, B [][]int) [][]int {
	rowsA, colsA := len(A), len(A[0])
	rowsB, colsB := len(B), len(B[0])

	if colsA != rowsB {
		panic("Matrix dimensions are incompatible for multiplication")
	}

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return result
}

func main() {
	var rowsA, colsA, rowsB, colsB int
	fmt.Scan(&rowsA, &colsA)

	A := make([][]int, rowsA)
	for i := range A {
		A[i] = make([]int, colsA)
		for j := range A[i] {
			fmt.Scan(&A[i][j])
		}
	}

	fmt.Scan(&rowsB, &colsB)

	B := make([][]int, rowsB)
	for i := range B {
		B[i] = make([]int, colsB)
		for j := range B[i] {
			fmt.Scan(&B[i][j])
		}
	}

	result := matrixMultiplication(A, B)

	for i := range result {
		for j := range result[i] {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
}
