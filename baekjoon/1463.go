package main

import (
	"fmt"
	"math"
)

func minOperations(n int) int {
	if n == 1 {
		return 0
	}

	dp := make([]int, n+1)
	dp[1] = 0

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i%2 == 0 {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i/2]+1)))
		}
		if i%3 == 0 {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i/3]+1)))
		}
	}

	return dp[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(minOperations(n))
}
