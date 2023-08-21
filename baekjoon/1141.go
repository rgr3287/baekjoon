package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func isPrefix(word, prefix string) bool {
	return len(word) >= len(prefix) && word[:len(prefix)] == prefix
}

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	words := make([]string, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		words[i] = scanner.Text()
	}

	// 단어들을 길이 순으로 정렬
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	// 접두사 검사를 통해 접두사X 집합의 크기 계산
	prefixFreeCount := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPrefix(words[j], words[i]) {
				prefixFreeCount--
				break
			}
		}
	}

	fmt.Println(prefixFreeCount)
}
