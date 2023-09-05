package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	From, To, Weight int
}

type UnionFind struct {
	Parent []int
	Size   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	sizeArr := make([]int, size)
	for i := range parent {
		parent[i] = i
		sizeArr[i] = 1
	}
	return &UnionFind{Parent: parent, Size: sizeArr}
}

func (uf *UnionFind) Find(x int) int {
	if uf.Parent[x] == x {
		return x
	}
	uf.Parent[x] = uf.Find(uf.Parent[x])
	return uf.Parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		if uf.Size[rootX] < uf.Size[rootY] {
			rootX, rootY = rootY, rootX
		}
		uf.Parent[rootY] = rootX
		uf.Size[rootX] += uf.Size[rootY]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	V, E := 0, 0
	fmt.Scanf("%d %d\n", &V, &E)

	edges := make([]Edge, E)

	for i := 0; i < E; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		A, _ := strconv.Atoi(parts[0])
		B, _ := strconv.Atoi(parts[1])
		C, _ := strconv.Atoi(parts[2])
		edges[i] = Edge{A, B, C}
	}

	// 간선을 가중치 순으로 정렬
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	uf := NewUnionFind(V)
	totalWeight := 0

	// 최소 스패닝 트리 구하기
	for _, edge := range edges {
		if uf.Find(edge.From-1) != uf.Find(edge.To-1) {
			uf.Union(edge.From-1, edge.To-1)
			totalWeight += edge.Weight
		}
	}

	fmt.Println(totalWeight)
}
