package main

import (
	"fmt"
)

func main() {
	var testCases int
	fmt.Scan(&testCases)

	for t := 0; t < testCases; t++ {
		var sequence string
		fmt.Scan(&sequence)

		counts := [8]int{}

		for i := 0; i < len(sequence)-2; i++ {
			substr := sequence[i : i+3]
			switch substr {
			case "TTT":
				counts[0]++
			case "TTH":
				counts[1]++
			case "THT":
				counts[2]++
			case "THH":
				counts[3]++
			case "HTT":
				counts[4]++
			case "HTH":
				counts[5]++
			case "HHT":
				counts[6]++
			case "HHH":
				counts[7]++
			}
		}

		fmt.Println(counts[0], counts[1], counts[2], counts[3], counts[4], counts[5], counts[6], counts[7])
	}
}
