package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	N, _ := nextInt(scanner)
	a := make([][]int, N+1)

	for i := 0; i <= N; i++ {
		a = append(a, []int{})
	}

	indegree := make([]int, N+1)
	times := make([]int, N+1)

	for i := 1; i <= N; i++ {
		times[i], _ = nextInt(scanner)
		for {
			num, _ := nextInt(scanner)
			if num == -1 {
				break
			}
			a[num] = append(a[num], i)
			indegree[i]++
		}
	}

	ans := topologicalSort(a, indegree, times, N)
	fmt.Fprintln(writer, ans)
}

func nextInt(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	token := scanner.Text()
	return strconv.Atoi(token)
}

func topologicalSort(a [][]int, indegree []int, times []int, N int) string {
	q := []int{}
	sb := ""

	for i := 1; i <= N; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	result := make([]int, N+1)

	for len(q) > 0 {
		now := q[0]
		q = q[1:]

		for _, next := range a[now] {
			indegree[next]--
			result[next] = max(result[next], result[now]+times[now])

			if indegree[next] == 0 {
				q = append(q, next)
			}
		}
	}

	for i := 1; i <= N; i++ {
		sb += fmt.Sprintf("%d\n", result[i]+times[i])
	}

	return sb
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
