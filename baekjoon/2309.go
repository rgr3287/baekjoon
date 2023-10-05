package main

import (
	"fmt"
	"sort"
)

func main() {
	const n = 9
	var dwarfs [n]int

	for i := 0; i < n; i++ {
		fmt.Scan(&dwarfs[i])
	}

	sort.Ints(dwarfs[:])

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				if k != i && k != j {
					sum += dwarfs[k]
				}
			}
			if sum == 100 {
				for k := 0; k < n; k++ {
					if k != i && k != j {
						fmt.Println(dwarfs[k])
					}
				}
				return
			}
		}
	}
}
