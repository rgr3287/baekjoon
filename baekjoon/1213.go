package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	alpha := make([]int, 26)

	for _, char := range line {
		if char >= 'A' && char <= 'Z' {
			idx := int(char - 'A')
			alpha[idx]++
		}
	}

	isOne := 0
	for _, count := range alpha {
		if count%2 != 0 {
			isOne++
		}
	}

	var answer string
	var sb strings.Builder

	if isOne > 1 {
		answer = "I'm Sorry Hansoo"
	} else {
		for i, count := range alpha {
			for r := 0; r < count/2; r++ {
				sb.WriteByte(byte(i + 'A'))
			}
		}

		answer = sb.String()
		end := reverseStringData(answer)

		sb.Reset()
		for i, count := range alpha {
			if count%2 == 1 {
				sb.WriteByte(byte(i + 'A'))
			}
		}
		answer += sb.String() + end
	}

	fmt.Println(answer)
}

func reverseStringData(s string) string {
	var result strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		result.WriteByte(s[i])
	}
	return result.String()
}
