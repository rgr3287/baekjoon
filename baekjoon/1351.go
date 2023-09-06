package main

import "fmt"

func main() {
	var N, P, Q int64
	fmt.Scan(&N, &P, &Q)

	memo := make(map[int64]int64)
	fmt.Println(computeAN(N, P, Q, memo))
}

func computeAN(N, P, Q int64, memo map[int64]int64) int64 {
	if N == 0 {
		return 1
	}

	if val, found := memo[N]; found {
		return val
	}

	result := computeAN(N/P, P, Q, memo) + computeAN(N/Q, P, Q, memo)
	memo[N] = result

	return result
}
