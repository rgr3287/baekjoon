package main

//
//var (
//	mapGrid   [][]int
//	N, M, min int
//	visited   [][]bool
//	nodes     [][]int
//)
//
//func dfs(x, y, n int) {
//	nodes = append(nodes, []int{x, y})
//	dx, dy := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
//	for i := 0; i < 4; i++ {
//		nx, ny := x+dx[i], y+dy[i]
//		if mapGrid[nx][ny] != 0 {
//			if mapGrid[nx][ny] != n {
//				min = minInt(min, mapGrid[nx][ny])
//				continue
//			}
//			if visited[nx][ny] {
//				continue
//			}
//			visited[nx][ny] = true
//			dfs(nx, ny, n)
//		} else {
//			min = 0
//		}
//	}
//}
//
//func minInt(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//
//	input, _ := reader.ReadString('\n')
//	input = strings.TrimSuffix(input, "\n")
//	tokens := strings.Split(input, " ")
//	N, _ = strconv.Atoi(tokens[0])
//	M, _ = strconv.Atoi(tokens[1])
//
//	mapGrid = make([][]int, N+2)
//	visited = make([][]bool, N+2)
//	for i := range mapGrid {
//		mapGrid[i] = make([]int, M+2)
//		visited[i] = make([]bool, M+2)
//	}
//
//	result := 0
//
//	for i := 1; i < N+1; i++ {
//		mapGrid[i][0] = 0
//		mapGrid[i][M+1] = 0
//		input, _ := reader.ReadString('\n')
//		input = strings.TrimSuffix(input, "\n")
//		for j := 1; j < M+1; j++ {
//			mapGrid[i][j] = int(input[j-1] - '0')
//		}
//	}
//
//	for i := 0; i < M+2; i++ {
//		mapGrid[0][i] = 0
//		mapGrid[N+1][i] = 0
//	}
//
//	for i := 1; i < 10; i++ {
//		for j := 1; j < N+1; j++ {
//			for k := 1; k < M+1; k++ {
//				if mapGrid[j][k] == i && !visited[j][k] {
//					nodes = [][]int{}
//					min = 9
//					visited[j][k] = true
//					dfs(j, k, i)
//					if min > i {
//						temp := min - i
//						for _, pos := range nodes {
//							mapGrid[pos[0]][pos[1]] = min
//							visited[pos[0]][pos[1]] = false
//							result += temp
//						}
//					}
//				}
//			}
//		}
//	}
//
//	fmt.Println(result)
//}
