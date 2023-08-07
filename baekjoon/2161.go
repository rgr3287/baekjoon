package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	queue := make([]int, n)
	for i := 0; i < n; i++ {
		queue[i] = i + 1
	}

	for len(queue) > 1 {
		writer.WriteString(fmt.Sprintf("%d ", queue[0]))
		queue = queue[1:]

		// Move the front element to the back
		queue = append(queue, queue[0])
		queue = queue[1:]
	}

	writer.WriteString(fmt.Sprintf("%d\n", queue[0]))
}
