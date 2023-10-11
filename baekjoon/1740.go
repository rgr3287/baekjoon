package main

import "fmt"

func main() {
	var N uint64
	fmt.Scan(&N)

	result := findNthSumOfDistinctPowersOf3(N)
	fmt.Println(result)
}

func findNthSumOfDistinctPowersOf3(N uint64) uint64 {
	power := uint64(1)
	sum := uint64(0)

	for N > 0 {
		if N%2 == 1 {
			sum += power
		}
		N /= 2
		power *= 3
	}

	return sum
}
