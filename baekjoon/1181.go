package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var words = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &words[i])
	}

	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) < len(words[j]) {
			return true
		} else if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		} else {
			return false
		}
	})

	for i := 0; i < n; i++ {
		if i > 0 && words[i-1] == words[i] {
			continue
		}
		fmt.Fprintln(writer, words[i])
	}
}
