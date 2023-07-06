package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	board := make([]string, 8)
	for i := 0; i < 8; i++ {
		scanner.Scan()
		board[i] = scanner.Text()
	}

	count := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 && board[i][j] == 'F' {
				count++
			}
		}
	}

	fmt.Println(count)
}
