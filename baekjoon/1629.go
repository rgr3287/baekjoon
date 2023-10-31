package main

import (
	"fmt"
)

var a, b, c int

func main() {
	_, err := fmt.Scanf("%d %d %d", &a, &b, &c)
	if err != nil {
		panic(err)
	}

	result := pow(a, b, c)
	fmt.Println(result)
}

func pow(base, exponent, modulus int) int {
	if base == 1 || exponent == 1 {
		return base % modulus
	}

	tmp := pow(base, exponent/2, modulus)

	if exponent%2 == 1 {
		return (tmp * tmp % modulus) * base % modulus
	}

	return tmp * tmp % modulus
}
