package main

import "fmt"

func main() {
	var N, area int
	fmt.Scan(&N)

	grid := make([][]bool, 100)
	for i := 0; i < 100; i++ {
		grid[i] = make([]bool, 100)
	}

	for i := 0; i < N; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		for j := x; j < x+10; j++ {
			for k := y; k < y+10; k++ {
				if !grid[j][k] {
					grid[j][k] = true
					area++
				}
			}
		}
	}

	fmt.Println(area)
}
