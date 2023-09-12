package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Graph struct {
	Vertices int
	Edges    int
	AdjList  map[int][]int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.AdjList[u] = append(g.AdjList[u], v)
	g.AdjList[v] = append(g.AdjList[v], u)
}

func (g *Graph) SortAdjList() {
	for key := range g.AdjList {
		sort.Ints(g.AdjList[key])
	}
}

func DFS(graph *Graph, start int, visited map[int]bool) {
	visited[start] = true
	fmt.Printf("%d ", start)

	for _, neighbor := range graph.AdjList[start] {
		if !visited[neighbor] {
			DFS(graph, neighbor, visited)
		}
	}
}

func BFS(graph *Graph, start int) {
	queue := []int{start}
	visited := make(map[int]bool)
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node)

		for _, neighbor := range graph.AdjList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	V, _ := strconv.Atoi(scanner.Text())

	graph := NewGraph(N)

	for i := 0; i < M; i++ {
		scanner.Scan()
		u, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		graph.AddEdge(u, v)
	}

	graph.SortAdjList()

	visited := make(map[int]bool)
	fmt.Printf("DFS 결과: ")
	DFS(graph, V, visited)
	fmt.Println()

	visited = make(map[int]bool)
	fmt.Printf("BFS 결과: ")
	BFS(graph, V)
	fmt.Println()
}
