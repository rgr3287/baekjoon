package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	for {
		fmt.Scan(&a)
		if a == 0 {
			break
		}
		fmt.Printf("%.2f\n", 1+a+math.Pow(a, 2)+math.Pow(a, 3)+math.Pow(a, 4))
	}
}
