package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	max := 0
	now := 0

	for i := 0; i < 4; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		values := strings.Fields(line)

		bye, _ := strconv.Atoi(values[0])
		hello, _ := strconv.Atoi(values[1])

		now = now - bye + hello
		max = getMax(max, now)
	}

	fmt.Println(max)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
