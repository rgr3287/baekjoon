package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(input[:len(input)-1])

	sum := 0
	for i := 0; i <= n; i++ {
		for j := i; j <= n; j++ {
			sum += i
			sum += j
		}
	}

	fmt.Println(sum)
}
