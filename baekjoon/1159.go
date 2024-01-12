package main

import (
	"bufio"
	"fmt"
	"os"
)

const ASCII = 97

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n := parseInteger(scanner.Text())

	alphabet := make([]int, 26)

	// 성씨 개수 파악
	for i := 0; i < n; i++ {
		scanner.Scan()
		name := scanner.Text()
		first := name[0]
		alphabet[first-ASCII]++
	}

	var result string

	// 5개 이상 비교
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i] >= 5 {
			result += string(i + ASCII)
		}
	}

	if result == "" {
		fmt.Print("PREDAJA")
	} else {
		fmt.Print(result)
	}
}

func parseInteger(s string) int {
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
}
