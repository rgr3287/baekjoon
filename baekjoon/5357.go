package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dedupe(dupString string) string {
	var dedupeString strings.Builder
	dedupeString.WriteString(string(dupString[0]))

	for i := 1; i < len(dupString); i++ {
		if dupString[i] != dupString[i-1] {
			dedupeString.WriteByte(dupString[i])
		}
	}

	return dedupeString.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		dupString, _ := reader.ReadString('\n')
		deduped := dedupe(strings.TrimSpace(dupString))
		fmt.Println(deduped)
	}
}
