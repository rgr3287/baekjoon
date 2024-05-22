package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Fields(line)
		if len(coords) != 2 {
			continue
		}

		a, errA := strconv.ParseFloat(coords[0], 64)
		b, errB := strconv.ParseFloat(coords[1], 64)
		if errA != nil || errB != nil {
			continue
		}

		if a > 0 && b > 0 {
			fmt.Println("Q1")
		} else if a < 0 && b > 0 {
			fmt.Println("Q2")
		} else if a < 0 && b < 0 {
			fmt.Println("Q3")
		} else if a > 0 && b < 0 {
			fmt.Println("Q4")
		} else if a == 0 || b == 0 {
			fmt.Println("AXIS")
		}
	}
}
