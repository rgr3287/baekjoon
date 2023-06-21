package main

import (
	"fmt"
	"sort"
)

func main() {
	// Reading the input values
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	B := make([]int, N)

	// Reading the values of array A
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	// Reading the values of array B
	for i := 0; i < N; i++ {
		fmt.Scan(&B[i])
	}

	// Sorting arrays A and B in non-decreasing order
	sort.Ints(A)
	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})

	// Calculating the minimum sum of products
	minSum := 0
	for i := 0; i < N; i++ {
		minSum += A[i] * B[i]
	}

	// Printing the result
	fmt.Println(minSum)
}
