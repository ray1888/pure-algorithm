package minial_spaning_tree

import "sort"

type edge struct {
	Weight int
	Src    int
	Dest   int
}

type edges []*edge

func (e edges) Len() int { return len(e) }
func (e edges) Less(i, j int) bool {
	return e[i].Weight < e[j].Weight
}
func (e edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type subSet struct {
	Parent int
	Rank   int
}

type subSets []subSet

func find(set subSets, node int) int {
	if set[node].Parent == node {
		return node
	}
	return find(set, set[node].Parent)
}

func union(set *subSets, x, y int) {
	xroot := find(*set, x)
	rroot := find(*set, y)
	if (*set)[xroot].Rank < (*set)[rroot].Rank {
		(*set)[xroot].Parent = rroot
	} else if (*set)[xroot].Rank > (*set)[rroot].Rank {
		(*set)[rroot].Parent = xroot
	} else {
		(*set)[rroot].Parent = xroot
		(*set)[xroot].Rank++
	}
}

type graph struct {
	V       int   // vertices
	EdgeNum int   // edge number
	Edges   edges //
}

func createGrpah(v, e int) *graph {
	g := new(graph)
	g.V = v
	g.EdgeNum = e
	g.Edges = make(edges, 0)
	return g
}

func kruskalMST(g *graph, set subSets) int {
	// use union find to implement duplicate v
	i := 0
	e := 0

	result := make(edges, 0) // 用于存储最小生成树
	for e < g.V-1 && i < g.EdgeNum {
		next_edge := g.Edges[i]

		x := find(set, next_edge.Src)
		y := find(set, next_edge.Dest)

		// If including this edge does't cause cycle,
		// include it in result and increment the index
		// of result for next edge
		if x != y {
			result = append(result, next_edge)
			union(&set, x, y)
		}
		i++
		e++
	}
	sum := 0
	for _, item := range result {
		sum += item.Weight
	}
	return sum
}

func KruskalMST(v, e int, edgeInts [][]int) int {
	g := createGrpah(v, e)
	for _, item := range edgeInts {
		newItem := new(edge)
		newItem.Weight = item[0]
		newItem.Src = item[1]
		newItem.Dest = item[2]
		g.Edges = append(g.Edges, newItem)
	}
	set := make(subSets, v)
	for i := 0; i < v; i++ {
		set[i].Parent = i
		set[i].Rank = 0
	}
	sort.Sort(&g.Edges)
	return kruskalMST(g, set)
}
