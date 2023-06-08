package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findNthDigitAfterDecimal(A, B, N int) (int, error) {
	// Check if A, B, and N are within the specified size constraints
	if A < 1 || A > 100000 || B < 1 || B > 100000 || N < 1 || N > 1000000 {
		return 0, fmt.Errorf("input out of bounds")
	}

	// Perform division and convert the result to a string
	quotient := float64(A) / float64(B)
	quotientStr := strconv.FormatFloat(quotient, 'f', -1, 64)

	// Find the index of the decimal point in the quotient string
	decimalIndex := strings.Index(quotientStr, ".")

	// Check if the decimal point exists and if N is within bounds
	if decimalIndex == -1 || N > len(quotientStr)-decimalIndex-1 {
		return 0, fmt.Errorf("Nth digit not found")
	}

	// Extract the Nth digit after the decimal point
	nthDigit, _ := strconv.Atoi(string(quotientStr[decimalIndex+N+1]))

	return nthDigit, nil
}

func main() {
	var A, B, N int
	fmt.Scan(&A, &B, &N)

	nthDigit, err := findNthDigitAfterDecimal(A, B, N)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(nthDigit)
}
