package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	k, _ := new(big.Int).SetString(input[0], 10)
	l, _ := strconv.Atoi(input[1])

	primes := make([]bool, l+1)
	primes[1] = true

	for i := 2; i < l; i++ {
		if primes[i] {
			continue // 소수가 아닌(true) 수는 넘어가기
		}

		now := new(big.Int)
		now.SetInt64(int64(i))

		if new(big.Int).Mod(k, now).Cmp(big.NewInt(0)) == 0 {
			// p를 now로 나눈 나머지가 0이면
			fmt.Printf("BAD %s\n", now.String())
			return
		}

		for j := i + i; j <= l; j += i {
			// i를 제외한 i의 배수 모두 체크하기
			primes[j] = true
		}
	}
	fmt.Println("GOOD")
}
