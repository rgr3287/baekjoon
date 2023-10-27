package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	next int
	dist int
}

func find(start, end int, N int, list [][]Node) int {
	visit := make([]bool, N+1)
	queue := make([]Node, 0)
	queue = append(queue, Node{start, 0})
	visit[start] = true

	var dist int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.next == end {
			dist = node.dist
			break
		}

		for _, temp := range list[node.next] {
			if !visit[temp.next] {
				queue = append(queue, Node{temp.next, node.dist + temp.dist})
				visit[temp.next] = true
			}
		}
	}

	return dist
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	list := make([][]Node, N+1)
	for i := 1; i <= N; i++ {
		list[i] = make([]Node, 0)
	}

	for i := 0; i < N-1; i++ {
		scanner.Scan()
		start, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		end, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		cost, _ := strconv.Atoi(scanner.Text())

		list[start] = append(list[start], Node{end, cost})
		list[end] = append(list[end], Node{start, cost})
	}

	for i := 0; i < M; i++ {
		scanner.Scan()
		start, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		end, _ := strconv.Atoi(scanner.Text())
		fmt.Println(find(start, end, N, list))
	}
}
