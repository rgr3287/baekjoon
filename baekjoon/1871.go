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
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Split(input, "-")
		a, b := parts[0], parts[1]

		valueA := 0
		for j := 0; j < len(a); j++ {
			value := int(a[j]-'A') * int(pow(26, 2-j))
			valueA += value
		}

		valueB, _ := strconv.Atoi(b)

		if abs(valueA-valueB) <= 100 {
			fmt.Fprintln(writer, "nice")
		} else {
			fmt.Fprintln(writer, "not nice")
		}
	}
}

func pow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
