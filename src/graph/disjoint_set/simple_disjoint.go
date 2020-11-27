package disjoint_set

func find(subset *[]int, item int) int {
	if (*subset)[item] == -1 {
		return item
	}
	return find(subset, (*subset)[item])
}

func union(subset *[]int, x, y int) {
	(*subset)[x] = y
}

// the  easiset version of
func IsCycle(graph [][]int, v int) bool {
	subset := make([]int, v)
	for i := 0; i < v; i++ {
		subset[i] = -1
	}

	for i := 0; i < len(graph); i++ {
		x := find(&subset, graph[i][1])
		y := find(&subset, graph[i][2])
		if x == y {
			return true
		}
		union(&subset, x, y)
	}
	return false
}
