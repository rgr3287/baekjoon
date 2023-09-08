package main

import (
	"container/heap"
	"fmt"
)

type EdgeData struct {
	to, cost int
}

func dijkstra(graph [][]EdgeData, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = int(1e9)
	}
	dist[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{start, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		u := cur.value
		if dist[u] < cur.priority {
			continue
		}
		for _, e := range graph[u] {
			v, w := e.to, e.cost
			if alt := dist[u] + w; alt < dist[v] {
				dist[v] = alt
				heap.Push(&pq, &Item{v, alt})
			}
		}
	}
	return dist
}

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	var n, m, x int
	fmt.Scan(&n, &m, &x)

	graph := make([][]Edge, n)
	for i := 0; i < m; i++ {
		var start, end, cost int
		fmt.Scan(&start, &end, &cost)
		start--
		end--
		graph[start] = append(graph[start], EdgeData{end, cost})
	}

	maxTime := 0
	for i := 0; i < n; i++ {
		distToX := dijkstra(graph, i)
		distFromX := dijkstra(graph, x-1)
		totalTime := distToX[x-1] + distFromX[i]
		if totalTime > maxTime {
			maxTime = totalTime
		}
	}

	fmt.Println(maxTime)
}
