package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]int, 101)
	for i := 0; i < 101; i++ {
		grid[i] = make([]int, 101)
	}

	for i := 0; i < n; i++ {
		var x1, y1, x2, y2 int
		fmt.Scan(&x1, &y1, &x2, &y2)

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				grid[x][y]++
			}
		}
	}

	count := 0
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			if grid[i][j] > m {
				count++
			}
		}
	}

	fmt.Println(count)
}
