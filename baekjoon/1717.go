package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UnionFindSt struct {
	parent []int
}

func (uf *UnionFindSt) init(N int) {
	uf.parent = make([]int, N+1)
	for i := 1; i <= N; i++ {
		uf.parent[i] = i
	}
}

func (uf *UnionFindSt) findData(x int) int {
	if x == uf.parent[x] {
		return x
	}
	uf.parent[x] = uf.findData(uf.parent[x])
	return uf.parent[x]
}

func (uf *UnionFindSt) union(x, y int) {
	x = uf.findData(x)
	y = uf.findData(y)

	if x != y {
		if x < y {
			uf.parent[y] = x
		} else {
			uf.parent[x] = y
		}
	}
}

func (uf *UnionFindSt) isSameParent(x, y int) bool {
	return uf.findData(x) == uf.findData(y)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	uf := UnionFindSt{}
	uf.init(N)

	var output strings.Builder

	for i := 0; i < M; i++ {
		scanner.Scan()
		command, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())

		if command == 0 {
			uf.union(a, b)
		} else if command == 1 {
			if uf.isSameParent(a, b) {
				output.WriteString("YES\n")
			} else {
				output.WriteString("NO\n")
			}
		}
	}

	fmt.Print(output.String())
}
