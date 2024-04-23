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

	for {
		scanner.Scan()
		input := scanner.Text()

		numbers := strings.Fields(input)
		A, _ := strconv.Atoi(numbers[0])
		B, _ := strconv.Atoi(numbers[1])

		if A == 0 && B == 0 {
			break
		}

		fmt.Println(A + B)
	}
}
