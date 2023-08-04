package main

import "fmt"

func main() {
	var grid [101][101]bool
	var x1, y1, x2, y2 int
	var count int

	for i := 0; i < 4; i++ {
		fmt.Scan(&x1, &y1, &x2, &y2)
		for x := x1; x < x2; x++ {
			for y := y1; y < y2; y++ {
				if !grid[x][y] {
					grid[x][y] = true
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
