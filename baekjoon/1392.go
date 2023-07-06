package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	for i := 0; i < q; i++ {
		var t int
		fmt.Scan(&t)
		score := findScore(scores, t)
		fmt.Println(score)
	}
}

func findScore(scores []int, t int) int {
	time := 0
	for i, score := range scores {
		time += score
		if t < time {
			return i + 1
		}
	}
	return -1
}
