package main

import "fmt"

const MOON_WEIGHT_RATIO = 0.167

func main() {
	for {
		var weight float64
		fmt.Scanln(&weight)

		if weight < 0.00 {
			break
		}

		fmt.Printf("Objects weighing %.2f on Earth will weigh %.2f on the moon.\n", weight, weight*MOON_WEIGHT_RATIO)
	}
}
