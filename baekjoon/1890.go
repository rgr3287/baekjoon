package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	board := make([][]int, N)
	for i := 0; i < N; i++ {
		board[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&board[i][j])
		}
	}

	// dp[i][j]는 (i, j) 위치에서 목적지까지 갈 수 있는 경로
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}

	dp[0][0] = 1 // 시작 위치에는 1개의 경로

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if dp[i][j] == 0 || (i == N-1 && j == N-1) { // 도달할 수 없는 칸이거나 목적지인 경우
				continue
			}

			// 오른쪽으로 이동한 경우
			if j+board[i][j] < N {
				dp[i][j+board[i][j]] += dp[i][j]
			}

			// 아래로 이동한 경우
			if i+board[i][j] < N {
				dp[i+board[i][j]][j] += dp[i][j]
			}
		}
	}

	fmt.Println(dp[N-1][N-1])
}
