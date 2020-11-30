package maxflow

import "math"

func dfs(rGraph [][]int, source int, visited []bool) {
	visited[source] = true
	for i := 0; i < len(rGraph); i++ {
		if rGraph[source][i] != 0 && !visited[i] {
			dfs(rGraph, i, visited)
		}
	}
}

func mincut(graph [][]int, s, t int) [][]int {
	result := make([][]int, 0)
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
	}

	visited := make([]bool, len(graph))
	dfs(rGraph, s, visited)
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph); j++ {
			if visited[i] && !visited[j] && graph[i][j] != 0 {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}
