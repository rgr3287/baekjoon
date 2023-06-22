package main

import (
	"fmt"
)

func main() {
	var k int
	var encodedText string
	fmt.Scanln(&k)
	fmt.Scanln(&encodedText)

	decodedText := decodeText(encodedText, k)
	fmt.Println(decodedText)
}

func decodeText(encodedText string, numRows int) string {
	numCols := len(encodedText) / numRows
	decoded := make([][]rune, numRows)
	for i := range decoded {
		decoded[i] = make([]rune, numCols)
	}

	index := 0
	for col := 0; col < numCols; col++ {
		if col%2 == 0 {
			for row := 0; row < numRows; row++ {
				decoded[row][col] = rune(encodedText[index])
				index++
			}
		} else {
			for row := numRows - 1; row >= 0; row-- {
				decoded[row][col] = rune(encodedText[index])
				index++
			}
		}
	}

	var decodedText string
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			decodedText += string(decoded[row][col])
		}
	}

	return decodedText
}
