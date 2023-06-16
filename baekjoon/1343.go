package main

import (
	"fmt"
	"strings"
)

func main() {
	var board string
	fmt.Scan(&board)

	result := solvePolyomino(board)
	fmt.Println(result)
}

func solvePolyomino(board string) string {
	board = strings.Replace(board, "XXXX", "AAAA", -1)
	board = strings.Replace(board, "XX", "BB", -1)

	if strings.Contains(board, "X") {
		return "-1"
	}

	return board
}
