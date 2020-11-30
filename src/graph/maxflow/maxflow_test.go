package maxflow

import "testing"

func Test_fordFulkerson(t *testing.T) {
	type args struct {
		graph [][]int
		s     int
		t     int
	}
	tests := []struct {
		name string
		args args
		want int
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
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fordFulkerson(tt.args.graph, tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("fordFulkerson() = %v, want %v", got, tt.want)
			}
		})
	}
}
