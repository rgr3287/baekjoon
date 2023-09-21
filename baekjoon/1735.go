package main

import "fmt"

// 최대공약수를 구하는 함수
func gcd2(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var a1, b1, a2, b2 int
	fmt.Scan(&a1, &b1)
	fmt.Scan(&a2, &b2)

	// 두 분수의 합을 구함
	numerator := a1*b2 + a2*b1
	denominator := b1 * b2

	// 최대공약수를 구함
	gcdVal := gcd2(numerator, denominator)

	// 기약분수 형태로 출력
	numerator /= gcdVal
	denominator /= gcdVal

	fmt.Printf("%d %d\n", numerator, denominator)
}
