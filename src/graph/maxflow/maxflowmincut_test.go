package maxflow

import "testing"

func Test_mincut(t *testing.T) {
	type args struct {
		graph [][]int
		s     int
		t     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				graph: [][]int{
					{0, 16, 13, 0, 0, 0},
					{0, 0, 10, 12, 0, 0},
					{0, 4, 0, 0, 14, 0},
					{0, 0, 9, 0, 0, 20},
					{0, 0, 0, 7, 0, 4},
					{0, 0, 0, 0, 0, 0},
				},
				s: 0,
				t: 5,
			},
			want: [][]int{
				{1, 3},
				{4, 3},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mincut(tt.args.graph, tt.args.s, tt.args.t)
		})
	}
}
