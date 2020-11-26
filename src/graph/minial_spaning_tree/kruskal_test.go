package minial_spaning_tree

import (
	"sort"
	"testing"
)

// func Test_kruskalMST(t *testing.T) {
// 	type args struct {
// 		g *graph
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := kruskalMST(tt.args.g); got != tt.want {
// 				t.Errorf("kruskalMST() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestKruskalMST(t *testing.T) {
	type args struct {
		v        int
		e        int
		edgeInts [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple three node ",
			args: args{
				v: 3,
				e: 3,
				edgeInts: [][]int{
					{10, 0, 2},
					{1, 0, 1},
					{7, 1, 2},
				},
			},
			want: 8,
		},
		{
			name: "five node ",
			args: args{
				v: 6,
				e: 8,
				edgeInts: [][]int{
					{10, 0, 2},
					{1, 0, 1},
					{7, 1, 2},
					{5, 2, 4},
					{1, 2, 3},
					{6, 3, 4},
					{3, 4, 5},
					{7, 3, 5},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KruskalMST(tt.args.v, tt.args.e, tt.args.edgeInts); got != tt.want {
				t.Errorf("KruskalMST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEdgesSort(t *testing.T) {
	edgesList := make(edges, 0)

	edgesList = append(edgesList, &edge{Weight: 1, Src: 5, Dest: 7})
	edgesList = append(edgesList, &edge{Weight: 5, Src: 3, Dest: 1})
	edgesList = append(edgesList, &edge{Weight: 2, Src: 4, Dest: 2})
	edgesList = append(edgesList, &edge{Weight: 6, Src: 2, Dest: 3})

	sort.Sort(edgesList)
	var result = []int{1, 2, 5, 6}
	for index, item := range edgesList {
		if result[index] != item.Weight {
			t.Fatalf("sort seqences error")
		}
	}
}
