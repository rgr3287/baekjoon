package main

import (
	"fmt"
	"strconv"
)

var board [5][5]int
var visited map[string]bool
var dx = [4]int{1, -1, 0, 0}
var dy = [4]int{0, 0, 1, -1}

func dfs(x, y, depth int, current string) {
	if depth == 6 {
		visited[current] = true
		return
	}

	for i := 0; i < 4; i++ {
		newX := x + dx[i]
		newY := y + dy[i]

		if newX >= 0 && newX < 5 && newY >= 0 && newY < 5 {
			newNum := strconv.Itoa(board[newX][newY])
			dfs(newX, newY, depth+1, current+newNum)
		}
	}
}

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&board[i][j])
		}
	}

	visited = make(map[string]bool)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			num := strconv.Itoa(board[i][j])
			dfs(i, j, 1, num)
		}
	}

	fmt.Println(len(visited))
}
