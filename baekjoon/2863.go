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

	line, _ := reader.ReadString('\n')
	tokens := strings.Fields(line)
	A, _ := strconv.ParseFloat(tokens[0], 64)
	B, _ := strconv.ParseFloat(tokens[1], 64)

	line, _ = reader.ReadString('\n')
	tokens = strings.Fields(line)
	C, _ := strconv.ParseFloat(tokens[0], 64)
	D, _ := strconv.ParseFloat(tokens[1], 64)

	arr := make([]float64, 4)
	arr[0] = (A / C) + (B / D)
	arr[1] = (C / D) + (A / B)
	arr[2] = (D / B) + (C / A)
	arr[3] = (B / A) + (D / C)

	max := arr[0]
	N := 0

	for i := 1; i < 4; i++ {
		if max < arr[i] {
			max = arr[i]
			N = i
		}
	}

	fmt.Println(N)
}
