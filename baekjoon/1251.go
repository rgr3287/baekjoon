package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Scan(&s)

	result := make([]string, 0)

	for i := 1; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			reversed := reverseString(s[:i]) + reverseString(s[i:j]) + reverseString(s[j:])
			result = append(result, reversed)
		}
	}

	sort.Strings(result)

	fmt.Println(result[0])
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
