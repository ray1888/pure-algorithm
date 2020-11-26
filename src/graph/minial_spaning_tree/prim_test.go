package minial_spaning_tree

import (
	"math"
	"testing"
)

func TestPrimMst(t *testing.T) {
	type args struct {
		graph GraphMatrix
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				graph: GraphMatrix{
					{0, 1, 2},
					{1, 0, 3},
					{1, 2, 0},
				},
			},
			want: 2,
		},
		{
			args: args{
				graph: GraphMatrix{
					{0, 1, 10, math.MaxInt32, math.MaxInt32, math.MaxInt32},
					{1, 0, 7, math.MaxInt32, math.MaxInt32, math.MaxInt32},
					{10, 7, 0, 1, 5, math.MaxInt32},
					{math.MaxInt32, 1, math.MaxInt32, 0, 6, 7},
					{math.MaxInt32, math.MaxInt32, 5, 6, 0, 3},
					{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32, 3, 0},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimMst(tt.args.graph); got != tt.want {
				t.Errorf("PrimMst() = %v, want %v", got, tt.want)
			}
		})
	}
}
