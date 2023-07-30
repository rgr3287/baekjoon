package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findBallCup(M int, swaps [][]int) int {
	ballPosition := 1
	for _, swap := range swaps {
		if swap[0] == ballPosition {
			ballPosition = swap[1]
		} else if swap[1] == ballPosition {
			ballPosition = swap[0]
		}
	}

	return ballPosition
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	swaps := make([][]int, M)
	for i := 0; i < M; i++ {
		scanner.Scan()
		swapStr := strings.Split(scanner.Text(), " ")
		X, _ := strconv.Atoi(swapStr[0])
		Y, _ := strconv.Atoi(swapStr[1])
		swaps[i] = []int{X, Y}
	}

	output := findBallCup(M, swaps)
	fmt.Println(output)
}
