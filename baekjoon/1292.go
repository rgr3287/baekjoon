package main

import "fmt"

func calculateSum(A, B int) int {
	sum := 0
	currentNumber := 1
	count := 1

	for i := 1; i <= B; i++ {
		if count > currentNumber {
			currentNumber++
			count = 1
		}

		if A <= i && i <= B {
			sum += currentNumber
		}

		count++
	}

	return sum
}

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	result := calculateSum(A, B)
	fmt.Println(result)
}
