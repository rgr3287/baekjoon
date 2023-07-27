package main

import "fmt"

func isGummin(number int) bool {
	for number > 0 {
		digit := number % 10
		if digit != 4 && digit != 7 {
			return false
		}
		number /= 10
	}
	return true
}

func findLargestGummin(N int) int {
	for i := N; i >= 4; i-- {
		if isGummin(i) {
			return i
		}
	}
	return -1
}

func main() {
	var N int
	fmt.Scan(&N)

	largestGummin := findLargestGummin(N)
	fmt.Println(largestGummin)
}
