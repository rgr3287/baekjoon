package main

import "fmt"

func main() {
	// Read input values
	var E, S, M int
	fmt.Scan(&E, &S, &M)

	// Find the year
	year := findYear(E, S, M)
	fmt.Println(year)
}

func findYear(E, S, M int) int {
	year := 1
	e, s, m := 1, 1, 1

	for {
		// Check if the current values match the input values
		if e == E && s == S && m == M {
			break
		}

		// Increment the values
		e = (e % 15) + 1
		s = (s % 28) + 1
		m = (m % 19) + 1

		year++
	}

	return year
}
