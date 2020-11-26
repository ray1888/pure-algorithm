package minial_spaning_tree

import (
	"math"
)

type GraphMatrix [][]int

func getminKey(keys []int, mstSet map[int]bool) int {
	minValue := math.MaxInt32
	minIndex := math.MaxInt32

	for v := 0; v < len(keys); v++ {
		if !mstSet[v] && keys[v] < minValue {
			minValue = keys[v]
			minIndex = v
		}
	}
	return minIndex
}

func PrimMst(graph GraphMatrix) int {
	lengthV := len(graph)
	parent := make([]int, lengthV)
	keys := make([]int, lengthV)
	mstSet := make(map[int]bool)

	for i := 0; i < lengthV; i++ {
		keys[i] = math.MaxInt32
	}
	keys[0] = 0
	parent[0] = -1

	for count := 0; count < lengthV-1; count++ {
		u := getminKey(keys, mstSet)
		mstSet[u] = true
		for v := 0; v < lengthV; v++ {
			if graph[u][v] != 0 && !mstSet[v] && graph[u][v] < keys[v] {
				parent[v] = u
				keys[v] = graph[u][v]
			}
		}
	}
	return keys[lengthV-1]
}
