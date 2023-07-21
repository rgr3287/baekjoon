package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	sequence := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&sequence[i])
	}

	for i := 0; i < K; i++ {
		newSequence := make([]int, len(sequence)-1)
		for j := 0; j < len(sequence)-1; j++ {
			newSequence[j] = sequence[j+1] - sequence[j]
		}
		sequence = newSequence
	}

	result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sequence)), ","), "[]")
	fmt.Println(result)
}
