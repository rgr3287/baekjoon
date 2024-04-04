package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := 0
	fmt.Sscanf(scanner.Text(), "%d", &n)

	var sb strings.Builder
	for i := 0; i < n; i++ {
		scanner.Scan()
		cur := strings.ToLower(scanner.Text())
		sb.WriteString(cur)
		sb.WriteString("\n")
	}
	fmt.Print(sb.String())
}

func main() {
	solution()
}
