package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numBirds, _ := strconv.Atoi(scanner.Text())

	time := 0
	currentNumber := 1

	for numBirds > 0 {
		if numBirds < currentNumber {
			currentNumber = 1
		}
		numBirds -= currentNumber
		time++
		currentNumber++
	}

	fmt.Println(time)
}
