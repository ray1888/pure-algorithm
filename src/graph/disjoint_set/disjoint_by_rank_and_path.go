package disjoint_set

type subSet struct {
	Parent int
	Rank   int
}

type subSets []subSet

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

type graph struct {
	V       int   // vertices
	EdgeNum int   // edge number
	Edges   edges //
}

func findAdv(set subSets, node int) int {
	if set[node].Parent == node {
		return node
	}
	return findAdv(set, set[node].Parent)
}

func unionAdv(set *subSets, x, y int) {
	xroot := findAdv((*set), x)
	rroot := findAdv((*set), y)
	if (*set)[xroot].Rank < (*set)[rroot].Rank {
		(*set)[xroot].Parent = rroot
	} else if (*set)[xroot].Rank > (*set)[rroot].Rank {
		(*set)[rroot].Parent = xroot
	} else {
		(*set)[rroot].Parent = xroot
		(*set)[xroot].Rank++
	}
}

func hasCycle(g *graph, set subSets) bool {
	i := 0
	e := 0
	cycleFlag := false

	for e < g.V && i < g.EdgeNum {
		next_edge := g.Edges[i]
		i++
		x := findAdv(set, next_edge.Src)
		y := findAdv(set, next_edge.Dest)
		if x == y {
			cycleFlag = true
			break
		}
		unionAdv(&set, x, y)
		e++
	}
	return cycleFlag
}
