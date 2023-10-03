package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	stack := make([]int, 0)
	result := make([]string, 0)

	nextToPush := 1
	possible := true

	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())

		for nextToPush <= num {
			stack = append(stack, nextToPush)
			result = append(result, "+")
			nextToPush++
		}

		if stack[len(stack)-1] == num {
			stack = stack[:len(stack)-1]
			result = append(result, "-")
		} else {
			possible = false
			break
		}
	}

	if possible {
		writer := bufio.NewWriter(os.Stdout)
		defer writer.Flush()

		for _, op := range result {
			fmt.Fprintln(writer, op)
		}
	} else {
		fmt.Println("NO")
	}
}
