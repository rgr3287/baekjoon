package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tree, lazy []int
var n int

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	tokens := strings.Split(input, " ")

	n, _ = strconv.Atoi(tokens[0])
	m, _ := strconv.Atoi(tokens[1])

	tree = make([]int, 4*n)
	lazy = make([]int, 4*n)

	var sb strings.Builder
	for i := 0; i < m; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		tokens := strings.Split(input, " ")

		op, _ := strconv.Atoi(tokens[0])
		a, _ := strconv.Atoi(tokens[1])
		b, _ := strconv.Atoi(tokens[2])

		if op == 0 {
			update(1, n, 1, a, b)
		} else {
			sb.WriteString(strconv.Itoa(pSum(1, n, 1, a, b)))
			sb.WriteString("\n")
		}
	}
	fmt.Print(sb.String())
}

func propagate(s, e, node int) {
	if lazy[node]%2 == 1 {
		if s != e {
			lazy[node*2] += lazy[node]
			lazy[node*2+1] += lazy[node]
		}
		tree[node] = (e - s + 1) - tree[node]
		lazy[node] = 0
	}
}

func update(s, e, node, l, r int) {
	propagate(s, e, node)
	if e < l || r < s {
		return
	}
	if l <= s && e <= r {
		lazy[node] = 1
		propagate(s, e, node)
		return
	}

	mid := (s + e) / 2
	update(s, mid, node*2, l, r)
	update(mid+1, e, node*2+1, l, r)
	tree[node] = tree[node*2] + tree[node*2+1]
}

func pSum(s, e, node, l, r int) int {
	propagate(s, e, node)
	if e < l || r < s {
		return 0
	}
	if l <= s && e <= r {
		return tree[node]
	}

	mid := (s + e) / 2
	return pSum(s, mid, node*2, l, r) + pSum(mid+1, e, node*2+1, l, r)
}
