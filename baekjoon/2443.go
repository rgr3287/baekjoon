package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	for i := N; i >= 1; i-- {
		for j := 0; j < N-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
