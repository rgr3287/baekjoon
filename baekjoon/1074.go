package main

import "fmt"

func main() {
	var N, r, c int
	fmt.Scan(&N, &r, &c)

	result := visitZ(N, r, c)
	fmt.Println(result)
}

func visitZ(N, r, c int) int {
	if N == 0 {
		return 0
	}

	halfSize := 1 << (N - 1)
	halfArea := halfSize * halfSize
	quad := 0

	if r >= halfSize {
		quad += 2
		r -= halfSize
	}

	if c >= halfSize {
		quad += 1
		c -= halfSize
	}

	return quad*halfArea + visitZ(N-1, r, c)
}
