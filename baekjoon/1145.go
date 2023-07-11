package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	numbers := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&numbers[i])
	}

	minimum := 1 << 30

	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			for k := j + 1; k < 5; k++ {
				lcmValue := lcm(numbers[i], lcm(numbers[j], numbers[k]))
				if lcmValue < minimum {
					minimum = lcmValue
				}
			}
		}
	}

	fmt.Println(minimum)
}
