package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var list [][]int
var dp []int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	st := scanner.Text()
	tokenizer := NewStringTokenizer(st)

	list = make([][]int, n)
	var rt int

	dp = make([]int, n)

	for i := 0; i < n; i++ {
		list[i] = make([]int, 0)
		a, _ := tokenizer.NextTokenInt()

		if a == -1 {
			rt = i
		} else {
			list[a] = append(list[a], i)
		}
	}

	min := solveF(rt)
	fmt.Println(min)
}

func solveF(idx int) int {
	for _, nxt := range list[idx] {
		dp[nxt] = 1 + solveF(nxt)
	}

	sort.Slice(list[idx], func(i, j int) bool {
		return dp[list[idx][j]] < dp[list[idx][i]]
	})

	var res int
	for i := 0; i < len(list[idx]); i++ {
		num := list[idx][i]
		dp[num] += i
		if dp[num] > res {
			res = dp[num]
		}
	}
	return res
}

type StringTokenizer struct {
	tokens []string
	index  int
}

func NewStringTokenizer(input string) *StringTokenizer {
	tokens := make([]string, 0)
	for _, token := range splitF(input, ' ') {
		tokens = append(tokens, token)
	}
	return &StringTokenizer{tokens: tokens, index: 0}
}

func (st *StringTokenizer) NextToken() string {
	token := st.tokens[st.index]
	st.index++
	return token
}

func (st *StringTokenizer) NextTokenInt() (int, error) {
	token := st.NextToken()
	num, err := strconv.Atoi(token)
	return num, err
}

func splitF(s string, delimiter rune) []string {
	var tokens []string
	token := ""
	for _, char := range s {
		if char == delimiter {
			tokens = append(tokens, token)
			token = ""
		} else {
			token += string(char)
		}
	}
	tokens = append(tokens, token)
	return tokens
}
