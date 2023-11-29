package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N   int
	str string
	dx  = []int{-1, 0, 1, 0}
	dy  = []int{0, 1, 0, -1}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	N, _ = strconv.Atoi(readLine(reader))
	str = readLine(reader)

	mapSize := 101
	mapGrid := make([][]int, mapSize)
	for i := range mapGrid {
		mapGrid[i] = make([]int, mapSize)
	}

	x, y := 50, 50
	maxX, maxY, minX, minY := 50, 50, 50, 50
	mapGrid[x][y] = 1

	dir := 2
	for i := 0; i < N; i++ {
		switch str[i] {
		case 'L':
			dir -= 1
			if dir == -1 {
				dir = 3
			}
		case 'R':
			dir = (dir + 1) % 4
		case 'F':
			x += dx[dir]
			y += dy[dir]
			maxX = max(maxX, x)
			maxY = max(maxY, y)
			minX = min(minX, x)
			minY = min(minY, y)
			mapGrid[x][y] = 1
		}
	}

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if mapGrid[i][j] == 1 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func readLine(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')
	return str[:len(str)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
