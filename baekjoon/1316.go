package main

import (
	"fmt"
)

func main() {
	var N, count int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var word string
		fmt.Scan(&word)
		if isGroupWord(word) {
			count++
		}
	}

	fmt.Println(count)
}

func isGroupWord(word string) bool {
	seen := make(map[rune]bool)
	prevChar := rune(0)

	for _, char := range word {
		if seen[char] && prevChar != char {
			return false
		}
		seen[char] = true
		prevChar = char
	}

	return true
}
