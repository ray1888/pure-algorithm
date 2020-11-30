package maxflow

import "math"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

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
				parent[i] = u
				visited[i] = true
			}
		}
	}
	return visited[t]
}

func fordFulkerson(graph [][]int, s, t int) int {
	maxFlow := 0
	u := 0
	rGraph := make([][]int, len(graph))
	for i := 0; i < len(graph[0]); i++ {
		rGraph[i] = make([]int, len(graph[0]))
		for j := 0; j < len(graph[0]); j++ {
			rGraph[i][j] = graph[i][j]
		}
	}
	parent := make([]int, len(graph))
	for bfs(rGraph, s, t, parent) {
		pathFlow := math.MaxInt32
		for v := t; v != s; v = parent[v] {
			u = parent[v]
			pathFlow = min(pathFlow, rGraph[u][v])
		}

		for v := t; v != s; v = parent[v] {
			u = parent[v]
			rGraph[u][v] -= pathFlow
			rGraph[v][u] += pathFlow
		}
		maxFlow += pathFlow
	}
	return maxFlow
}
