package main

import (
	"fmt"
)

const div = 1000000007

func solve(n, l, r int) int {
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	dp[1][1][1] = 1
	for i := 2; i <= n; i++ {
		dp[i][i][1] = 1
		dp[i][1][i] = 1
		for j := 1; j <= l; j++ {
			for k := 1; k <= r; k++ {
				dp[i][j][k] = (dp[i-1][j][k-1] + dp[i-1][j-1][k] + dp[i-1][j][k]*(i-2)) % div
			}
		}
	}

	return dp[n][l][r]
}

func main() {
	var n, l, r int
	fmt.Scanf("%d %d %d", &n, &l, &r)

	result := solve(n, l, r)
	fmt.Println(result)
}
