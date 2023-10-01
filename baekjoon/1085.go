package main

import (
	"fmt"
)

func main() {
	var x, y, w, h int
	fmt.Scan(&x, &y, &w, &h)

	minX := minFc(x, w-x)
	minY := minFc(y, h-y)
	minDist := minFc(minX, minY)

	fmt.Println(minDist)
}

func minFc(a, b int) int {
	if a < b {
		return a
	}
	return b
}
