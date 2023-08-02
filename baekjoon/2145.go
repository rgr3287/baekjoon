package main

import (
	"fmt"
)

func sumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func main() {
	var n int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		for n >= 10 {
			n = sumOfDigits(n)
		}
		fmt.Println(n)
	}
}
