package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var list [][]int
var visit []bool
var leafCnt int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	list = make([][]int, N)
	visit = make([]bool, N)

	for i := 0; i < N; i++ {
		list[i] = make([]int, 0)
	}

	var root int
	scanner.Scan()
	tokens := strings.Split(scanner.Text(), " ")
	for i := 0; i < N; i++ {
		temp, _ := strconv.Atoi(tokens[i])
		if temp == -1 {
			root = i
		} else {
			list[i] = append(list[i], temp)
			list[temp] = append(list[temp], i)
		}
	}

	scanner.Scan()
	delNode, _ := strconv.Atoi(scanner.Text())
	if delNode == root {
		fmt.Println(0)
		return
	}
	visit[delNode] = true

	dfs3(root)
	fmt.Println(leafCnt)
}

func dfs3(root int) {
	isLeaf := true
	visit[root] = true

	for i := 0; i < len(list[root]); i++ {
		if !visit[list[root][i]] {
			isLeaf = false
			dfs3(list[root][i])
		}
	}

	if isLeaf {
		leafCnt++
	}
}
