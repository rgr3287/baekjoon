package main

import (
	"fmt"
)

func countTrailingZeros(n int) int {
	count := 0
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}

func main() {
	var N int
	fmt.Scan(&N)

	result := countTrailingZeros(N)
	fmt.Println(result)
}
