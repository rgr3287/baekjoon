package main

import (
	"fmt"
	"math"
)

func main() {
	var numTiles int
	fmt.Scan(&numTiles)

	sideLength := int(math.Sqrt(float64(numTiles)))
	fmt.Printf("The largest square has side length %d.\n", sideLength)
}
