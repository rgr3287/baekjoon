package main

import (
	"fmt"
	"math"
)

const INF = math.MaxInt32

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	// Initialize the adjacency matrix
	adj := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if i == j {
				adj[i][j] = 0
			} else {
				adj[i][j] = INF
			}
		}
	}

	// Read the input edges
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		adj[u][v] = 1
		adj[v][u] = 1
	}

	// Perform Floyd-Warshall algorithm
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				adj[i][j] = min(adj[i][j], adj[i][k]+adj[k][j])
			}
		}
	}

	// Find the person with the minimum Bacon number
	minBaconNum := INF
	minPerson := 0
	for i := 1; i <= n; i++ {
		sum := 0
		for j := 1; j <= n; j++ {
			sum += adj[i][j]
		}
		if sum < minBaconNum {
			minBaconNum = sum
			minPerson = i
		}
	}

	fmt.Println(minPerson)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
