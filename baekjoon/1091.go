package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	N          int
	P, S       []int
	card, temp []int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	N, _ = strconv.Atoi(readLine(reader))
	P = make([]int, N)
	S = make([]int, N)
	card = make([]int, N)
	temp = make([]int, N)
	cnt := 0

	for i := 0; i < N; i++ {
		card[i] = i % 3
	}

	PInput := readLine(reader)
	SInput := readLine(reader)

	PStrings := strings.Fields(PInput)
	for i := 0; i < N; i++ {
		P[i], _ = strconv.Atoi(PStrings[i])
	}

	SStrings := strings.Fields(SInput)
	for i := 0; i < N; i++ {
		S[i], _ = strconv.Atoi(SStrings[i])
	}

	for {
		isMatched := true
		inf := true

		for i := 0; i < N; i++ {
			if card[i] != P[i] {
				isMatched = false
				break
			}
		}

		if !isMatched {
			for i := 0; i < N; i++ {
				temp[i] = card[S[i]]
			}
			copy(card, temp)
			cnt++
		} else {
			fmt.Println(cnt)
			break
		}

		for i := 0; i < N; i++ {
			if card[i] != i%3 {
				inf = false
			}
		}

		if inf {
			fmt.Println(-1)
			break
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')
	return strings.TrimSpace(str)
}
