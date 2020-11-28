package maxflow

//
func bfs(rGraph [][]int, s, t int, parent []int) bool {
	V := len(rGraph)
	visited := make([]bool, V)

	queue := make([]int, 0)
	queue = append(queue, s)
	// because s is starting point , can think as root
	visited[s] = true
	parent[s] = -1

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for i := 0; i < V; i++ {
			if !visited[i] && rGraph[u][i] > 0 {
				queue = append(queue, i)
			}
		}
	}

}

func fordFulkerson(graph [][]int, s, t int) int {
	maxFlow := 0

	return maxFlow
}
