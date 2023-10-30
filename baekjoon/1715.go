package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	pq := &PrioritiesQueue{}
	heap.Init(pq)

	for i := 0; i < N; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		heap.Push(pq, num)
	}

	sum := 0
	for pq.Len() > 1 {
		first := heap.Pop(pq).(int)
		twice := heap.Pop(pq).(int)
		sum += first + twice
		heap.Push(pq, first+twice)
	}

	fmt.Println(sum)
}

type PrioritiesQueue []int

func (pq PrioritiesQueue) Len() int           { return len(pq) }
func (pq PrioritiesQueue) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq PrioritiesQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PrioritiesQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PrioritiesQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
