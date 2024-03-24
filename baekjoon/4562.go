package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var sb strings.Builder
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		if x < y {
			sb.WriteString("NO BRAINS\n")
		} else {
			sb.WriteString("MMM BRAINS\n")
		}
	}
	fmt.Print(sb.String())
}
