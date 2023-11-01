package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var N, K int
var visitedd [100001]int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()
	inputs := split(input)

	N, _ = strconv.Atoi(inputs[0])
	K, _ = strconv.Atoi(inputs[1])

	result := bfs(N)
	fmt.Println(result)
}

func split(input string) []string {
	return strings.Fields(input)
}

func bfs(node int) int {
	queue := make([]int, 0)

	queue = append(queue, node)
	index := node
	var n int
	visitedd[index] = 1

	for len(queue) > 0 {
		n, queue = queue[0], queue[1:]

		if n == K {
			return visitedd[n] - 1
		}

		if n-1 >= 0 && visitedd[n-1] == 0 {
			visitedd[n-1] = visitedd[n] + 1
			queue = append(queue, n-1)
		}
		if n+1 <= 100000 && visitedd[n+1] == 0 {
			visitedd[n+1] = visitedd[n] + 1
			queue = append(queue, n+1)
		}
		if 2*n <= 100000 && visitedd[2*n] == 0 {
			visitedd[2*n] = visitedd[n] + 1
			queue = append(queue, 2*n)
		}
	}
	return -1
}
