package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	floor := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&floor[i])
	}

	horizontalCount := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if floor[i][j] == '-' {
				horizontalCount++
				for j+1 < m && floor[i][j+1] == '-' {
					j++
				}
			}
		}
	}

	verticalCount := 0
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if floor[i][j] == '|' {
				verticalCount++
				for i+1 < n && floor[i+1][j] == '|' {
					i++
				}
			}
		}
	}

	totalDecorations := horizontalCount + verticalCount
	fmt.Println(totalDecorations)
}
