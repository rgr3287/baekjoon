package main

import (
	"bufio"
	"fmt"
	"os"
)

func footballTeam(name string) string {
	changeWordDict := map[rune]rune{'e': 'i', 'i': 'e', 'E': 'I', 'I': 'E'}

	var newName []rune
	for _, char := range name {
		if newChar, ok := changeWordDict[char]; ok {
			newName = append(newName, newChar)
		} else {
			newName = append(newName, char)
		}
	}

	return string(newName)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		name := scanner.Text()
		fmt.Println(footballTeam(name))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
