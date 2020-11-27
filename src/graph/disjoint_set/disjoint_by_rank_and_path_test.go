package disjoint_set

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		g   *graph
		set subSets
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has cycle",
			args: args{
				g: &graph{
					V:       3,
					EdgeNum: 3,
					Edges: edges{
						{Weight: 10, Src: 0, Dest: 2},
						{Weight: 10, Src: 0, Dest: 1},
						{Weight: 20, Src: 1, Dest: 2},
					},
				},
				set: subSets{
					{Parent: 0, Rank: 0},
					{Parent: 1, Rank: 0},
					{Parent: 2, Rank: 0},
				},
			},
			want: true,
		},
		{
			name: "has not  cycle",
			args: args{
				g: &graph{
					V:       3,
					EdgeNum: 2,
					Edges: edges{
						{Weight: 10, Src: 0, Dest: 2},
						{Weight: 10, Src: 0, Dest: 1},
					},
				},
				set: subSets{
					{Parent: 0, Rank: 0},
					{Parent: 1, Rank: 0},
					{Parent: 2, Rank: 0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.g, tt.args.set); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
