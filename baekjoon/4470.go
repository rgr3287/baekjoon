package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Printf("%d. %s\n", i+1, scanner.Text())
	}
}
