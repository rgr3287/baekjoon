package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	i := 1

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		text = strings.TrimSpace(text)
		inputs := strings.Fields(text)

		A, _ := strconv.ParseFloat(inputs[0], 64)
		B, _ := strconv.ParseFloat(inputs[1], 64)
		C, _ := strconv.ParseFloat(inputs[2], 64)

		if B == 0 {
			break
		}

		C = (C / 60) / 60
		distance := ((((A * math.Pi) * B) / 12) / 5280)
		MPH := distance / C

		fmt.Printf("Trip #%d: %.2f %.2f\n", i, distance, MPH)
		i++
	}
}
