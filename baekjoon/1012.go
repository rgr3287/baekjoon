package main

import (
	"fmt"
)

var (
	dx = [4]int{-1, 1, 0, 0} // 상하좌우 방향을 나타내는 배열
	dy = [4]int{0, 0, -1, 1}
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var M, N, K int
		fmt.Scan(&M, &N, &K)

		field := make([][]bool, M)
		for i := range field {
			field[i] = make([]bool, N)
		}

		for i := 0; i < K; i++ {
			var X, Y int
			fmt.Scan(&X, &Y)
			field[X][Y] = true
		}

		count := 0
		for i := 0; i < M; i++ {
			for j := 0; j < N; j++ {
				if field[i][j] {
					dfs(field, i, j)
					count++
				}
			}
		}

		fmt.Println(count)
	}
}

func dfs(field [][]bool, x, y int) {
	field[x][y] = false

	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx >= 0 && ny >= 0 && nx < len(field) && ny < len(field[0]) && field[nx][ny] {
			dfs(field, nx, ny)
		}
	}
}
