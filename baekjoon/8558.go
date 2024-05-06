package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		result %= 10
	}
	return result
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(factorial(n))
}
