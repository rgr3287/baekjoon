package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		var targetChar string
		fmt.Fscanf(reader, "%s ", &targetChar)

		if targetChar == "#" {
			break
		}

		sentence, _ := reader.ReadString('\n')
		sentence = strings.TrimSpace(sentence)

		count := 0
		for _, char := range sentence {
			if unicode.ToLower(char) == rune(targetChar[0]) {
				count++
			}
		}

		fmt.Fprintf(writer, "%s %d\n", targetChar, count)
	}
}
