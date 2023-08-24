package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var M, N int
	fmt.Scan(&M, &N)

	numbers := make([]int, N-M+1)
	for i := M; i <= N; i++ {
		numbers[i-M] = i
	}

	sort.Slice(numbers, func(i, j int) bool {
		return compareNumbers(numbers[i], numbers[j])
	})

	for i, num := range numbers {
		if i > 0 && i%10 == 0 {
			fmt.Println()
		} else if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()
}

func compareNumbers(a, b int) bool {
	englishA := numToEnglish(a)
	englishB := numToEnglish(b)
	return englishA < englishB
}

func numToEnglish(num int) string {
	englishNumbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digits := []int{}

	for num > 0 {
		digit := num % 10
		digits = append([]int{digit}, digits...)
		num /= 10
	}

	english := ""
	for _, digit := range digits {
		english += englishNumbers[digit] + " "
	}

	return strings.TrimSpace(english)
}
