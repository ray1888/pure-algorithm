package disjoint_set

import (
	"reflect"
	"testing"
)

func Test_disjoint(t *testing.T) {
	type args struct {
		graph [][]int
		v     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has cycle",
			args: args{
				// array item is weight, src, desc
				graph: [][]int{
					{1, 1, 0},
					{3, 0, 2},
					{1, 1, 2},
				},
				v: 3,
			},
			want: true,
		},
		{
			name: "has not cycle",
			args: args{
				// array item is weight, src, desc
				graph: [][]int{
					{1, 1, 0},
					{3, 2, 3},
					{1, 1, 2},
				},
				v: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCycle(tt.args.graph, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("disjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
