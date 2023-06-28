package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "*" {
			break
		}

		isSurprising := isStringSurprising(line)
		if isSurprising {
			fmt.Printf("%s is surprising.\n", line)
		} else {
			fmt.Printf("%s is NOT surprising.\n", line)
		}
	}
}

func isStringSurprising(str string) bool {
	n := len(str)

	// Check for each D-pair from 0 to (n-2)
	for d := 0; d <= n-2; d++ {
		pairs := make(map[string]bool)

		// Generate and check all D-pairs
		for i := 0; i <= n-2-d; i++ {
			pair := string(str[i]) + string(str[i+d+1])

			// If the pair already exists, the string is not surprising
			if pairs[pair] {
				return false
			}

			// Mark the pair as seen
			pairs[pair] = true
		}
	}

	// If all D-pairs are unique, the string is surprising
	return true
}
