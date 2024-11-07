package day1

import "testing"

func Test_findTargetIn2DPlants(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				matrix: [][]int{
					{1, 3, 5},
					{2, 5, 7},
				},
				target: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetIn2DPlants(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findTargetIn2DPlants() = %v, want %v", got, tt.want)
			}
		})
	}
}
