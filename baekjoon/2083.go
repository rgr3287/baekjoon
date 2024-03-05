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

	for {
		scanner.Scan()
		input := scanner.Text()
		values := strings.Fields(input)

		if len(values) == 3 {
			name := values[0]
			age, _ := strconv.Atoi(values[1])
			weight, _ := strconv.Atoi(values[2])

			if name == "#" && age == 0 && weight == 0 {
				break
			}

			if age > 17 || weight >= 80 {
				fmt.Println(name + " Senior")
			} else {
				fmt.Println(name + " Junior")
			}
		}
	}
}
