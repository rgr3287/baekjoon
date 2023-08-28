package main

import (
	"fmt"
)

func isHanNumber(n int) bool {
	if n < 100 {
		return true // 1부터 99까지는 모두 한수
	}

	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	difference := digits[0] - digits[1]
	for i := 1; i < len(digits)-1; i++ {
		if difference != digits[i]-digits[i+1] {
			return false
		}
	}

	return true
}

func main() {
	var N, count int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		if isHanNumber(i) {
			count++
		}
	}

	fmt.Println(count)
}
