package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	result := fibonacci(n)
	fmt.Println(result)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
