package main

import (
	"fmt"
)

func oddities(n int) string {
	if n % 2 == 0 {
		return fmt.Sprintf("%d is even", n)
	} else {
		return fmt.Sprintf("%d is odd", n)
	}
}

func main() {
	var testCases int
	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var n int
		fmt.Scan(&n)
		fmt.Println(oddities(n))
	}
}
