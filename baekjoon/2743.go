package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Scan(&word)
	fmt.Println(len(strings.TrimSpace(word)))
}
