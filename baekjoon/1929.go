package main

import (
	"fmt"
)

func sieveOfEratosthenes(n int) []bool {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	return isPrime
}

func main() {
	var M, N int
	fmt.Scan(&M, &N)

	isPrime := sieveOfEratosthenes(N)

	for i := M; i <= N; i++ {
		if isPrime[i] {
			fmt.Println(i)
		}
	}
}
