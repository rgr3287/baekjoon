package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Value int
	Index int
}

func main() {
	var scores [8]Score

	for i := 0; i < 8; i++ {
		var score int
		fmt.Scan(&score)
		scores[i] = Score{Value: score, Index: i + 1}
	}

	sort.Slice(scores[:], func(i, j int) bool {
		return scores[i].Value > scores[j].Value
	})

	totalSum := 0
	selectedIndices := []int{}

	for i := 0; i < 5; i++ {
		totalSum += scores[i].Value
		selectedIndices = append(selectedIndices, scores[i].Index)
	}

	sort.Ints(selectedIndices)

	fmt.Println(totalSum)
	for _, index := range selectedIndices {
		fmt.Printf("%d ", index)
	}
}
