package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	S1 := parseInt(scanner.Text())
	scanner.Scan()
	S2 := parseInt(scanner.Text())
	scanner.Scan()
	S3 := parseInt(scanner.Text())

	// Calculate the frequency of each sum
	sums := make([]int, S1+S2+S3+1) // Frequency of sums (indices represent the sum)
	maxFreq := 0                    // Maximum frequency
	maxSum := S1 + S2 + S3          // Sum with the highest frequency

	for i := 1; i <= S1; i++ {
		for j := 1; j <= S2; j++ {
			for k := 1; k <= S3; k++ {
				sum := i + j + k
				sums[sum]++
				if sums[sum] > maxFreq || (sums[sum] == maxFreq && sum < maxSum) {
					maxFreq = sums[sum]
					maxSum = sum
				}
			}
		}
	}

	// Print the sum with the highest frequency
	fmt.Println(maxSum)
}

// Helper function to convert a string to an integer
func parseInt(s string) int {
	var res int
	fmt.Sscanf(s, "%d", &res)
	return res
}
